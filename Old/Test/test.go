package main

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	helmv2 "github.com/fluxcd/helm-controller/api/v2beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var CoreComponents = map[string]struct {
	ComponentName string
	ComponentType string
	Step          int
}{
	"seaweedfs":       {ComponentName: "seaweedfs", ComponentType: "infra", Step: 2},
	"oci-registry":    {ComponentName: "sh-registry", ComponentType: "infra", Step: 3},
	"vault-operator":  {ComponentName: "vault-operator", ComponentType: "infra", Step: 4},
	"nats":            {ComponentName: "nats", ComponentType: "infra", Step: 5},
	"prometheus":      {ComponentName: "prometheus", ComponentType: "infra", Step: 6},
	"opentelemetry":   {ComponentName: "opentelemetry", ComponentType: "infra", Step: 7},
	"fluentbit":       {ComponentName: "fluentbit", ComponentType: "infra", Step: 8},
	"istio":           {ComponentName: "istio", ComponentType: "infra", Step: 9},
	"gloo-edge":       {ComponentName: "gloo-edge", ComponentType: "infra", Step: 10},
	"Runtime service": {ComponentName: "runtime2", ComponentType: "runtime-service", Step: 11},
}

type ComponentStatus struct {
	Name    string
	Title   string
	Status  string
	Message string
	Step    int
}

const (
	StatusCompleted  = "Completed"
	StatusInProgress = "InProgress"
	StatusFailed     = "Failed"
	StatusPending    = "Pending"
)

func main() {
	k8sClient, err := setupKubernetesClient()
	if err != nil {
		log.Fatalf("Error setting up Kubernetes client: %v", err)
	}

	fluxCDStatus := getFluxCDStatusWithRetry(k8sClient)
	fluxCDStatus.Step = 1

	var coreComponentsStatus []ComponentStatus
	if fluxCDStatus.Status == StatusFailed || fluxCDStatus.Status == StatusPending {
		coreComponentsStatus = setAllComponentsToPending(fluxCDStatus.Message)
	} else {
		coreComponentsStatus = fetchComponentsStatus(k8sClient)
	}

	printComponentStatuses(fluxCDStatus, coreComponentsStatus)
}

func setupKubernetesClient() (client.Client, error) {
	scheme := runtime.NewScheme()
	_ = appsv1.AddToScheme(scheme)
	_ = corev1.AddToScheme(scheme)
	_ = helmv2.AddToScheme(scheme)

	cfg, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("error getting kubeconfig: %v", err)
	}

	k8sClient, err := client.New(cfg, client.Options{Scheme: scheme})
	if err != nil {
		return nil, fmt.Errorf("error creating Kubernetes client: %v", err)
	}

	return k8sClient, nil
}

func getFluxCDStatusWithRetry(k8sClient client.Client) ComponentStatus {
	for i := 0; i < 3; i++ {
		fluxCDStatus := getFluxCDStatus(k8sClient)
		if fluxCDStatus.Status != StatusFailed {
			return fluxCDStatus
		}

		if i < 2 {
			log.Printf("FluxCD check attempt %d failed. Retrying in 2 seconds...", i+1)
			time.Sleep(2 * time.Second)
		}
	}

	return ComponentStatus{
		Name:    "FluxCD",
		Status:  StatusFailed,
		Message: "FluxCD failed to initialize after 3 attempts",
		Step:    1,
	}
}

func setAllComponentsToPending(reason string) []ComponentStatus {
	var components []ComponentStatus
	for componentName, compInfo := range CoreComponents {
		components = append(components, ComponentStatus{
			Name:    componentName,
			Title:   compInfo.ComponentType,
			Status:  StatusPending,
			Message: fmt.Sprintf("%s: Pending due to FluxCD status - %s", componentName, reason),
			Step:    compInfo.Step,
		})
	}
	return components
}

func fetchComponentsStatus(k8sClient client.Client) []ComponentStatus {
	var components []ComponentStatus
	helmReleases := listHelmReleases(k8sClient)

	for componentName, compInfo := range CoreComponents {
		componentStatus := processComponent(k8sClient, helmReleases, compInfo, componentName)
		components = append(components, componentStatus)
	}

	sort.Slice(components, func(i, j int) bool {
		return components[i].Step < components[j].Step
	})

	return components
}

func processComponent(k8sClient client.Client, helmReleases []helmv2.HelmRelease, compInfo struct {
	ComponentName string
	ComponentType string
	Step          int
}, componentName string) ComponentStatus {
	var associatedReleases []helmv2.HelmRelease
	for _, hr := range helmReleases {
		if hr.Labels["outsystems-component"] == "true" &&
			hr.Labels["component-name"] == compInfo.ComponentName &&
			hr.Labels["component-type"] == compInfo.ComponentType {
			associatedReleases = append(associatedReleases, hr)
		}
	}

	if len(associatedReleases) == 0 {
		return ComponentStatus{
			Name:    componentName,
			Title:   compInfo.ComponentType,
			Status:  StatusPending,
			Message: fmt.Sprintf("%s: Pending - Component not found or pending installation", componentName),
			Step:    compInfo.Step,
		}
	}

	status, message := aggregateComponentStatus(k8sClient, associatedReleases, componentName)

	return ComponentStatus{
		Name:    componentName,
		Title:   compInfo.ComponentType,
		Status:  status,
		Message: message,
		Step:    compInfo.Step,
	}
}

func aggregateComponentStatus(k8sClient client.Client, releases []helmv2.HelmRelease, componentName string) (string, string) {
	statusCounts := map[string]int{
		StatusCompleted:  0,
		StatusPending:    0,
		StatusInProgress: 0,
		StatusFailed:     0,
	}

	var messages []string

	for _, hr := range releases {
		status, message := getHelmReleaseStatus(k8sClient, hr)
		statusCounts[status]++
		messages = append(messages, fmt.Sprintf("%s: %s (%s)", hr.Name, message, status))
	}

	totalCount := len(releases)

	overallStatus := StatusInProgress
	if statusCounts[StatusFailed] > 0 {
		overallStatus = StatusFailed
	} else if statusCounts[StatusPending] == totalCount {
		overallStatus = StatusPending
	} else if statusCounts[StatusCompleted] == totalCount {
		overallStatus = StatusCompleted
	}

	detailedMessage := fmt.Sprintf("%s: %d/%d Completed, %d/%d Pending, %d/%d In Progress, %d/%d Failed. Details: %s",
		componentName,
		statusCounts[StatusCompleted], totalCount,
		statusCounts[StatusPending], totalCount,
		statusCounts[StatusInProgress], totalCount,
		statusCounts[StatusFailed], totalCount,
		strings.Join(messages, "; "))

	return overallStatus, detailedMessage
}

func getHelmReleaseStatus(k8sClient client.Client, hr helmv2.HelmRelease) (string, string) {
	for _, condition := range hr.Status.Conditions {
		if condition.Type == helmv2.ReleasedCondition {
			switch condition.Status {
			case metav1.ConditionTrue:
				return StatusCompleted, "Installed"
			case metav1.ConditionFalse:
				if condition.Reason == "InstallFailed" {
					return StatusFailed, getHelmReleaseError(k8sClient, hr)
				}
				return StatusInProgress, "Installing"
			}
		}
	}
	return StatusPending, "Pending installation"
}

func getHelmReleaseError(k8sClient client.Client, hr helmv2.HelmRelease) string {
	var updatedHR helmv2.HelmRelease
	err := k8sClient.Get(context.Background(), client.ObjectKey{Name: hr.Name, Namespace: hr.Namespace}, &updatedHR)
	if err != nil {
		return fmt.Sprintf("Error fetching HelmRelease details: %v", err)
	}

	for _, condition := range updatedHR.Status.Conditions {
		if condition.Type == helmv2.ReleasedCondition && condition.Status == metav1.ConditionFalse {
			return condition.Message
		}
	}

	return "Unknown error"
}

func getFluxCDStatus(k8sClient client.Client) ComponentStatus {
	fluxControllers := []string{
		"source-controller",
		"kustomize-controller",
		"helm-controller",
		"notification-controller",
		"image-reflector-controller",
		"image-automation-controller",
	}

	overallStatus := StatusCompleted
	var statusMessages []string

	for _, controllerName := range fluxControllers {
		deploymentStatus := checkControllerStatus(k8sClient, controllerName)
		if deploymentStatus.Status != StatusCompleted {
			overallStatus = deploymentStatus.Status
			statusMessages = append(statusMessages, fmt.Sprintf("%s: %s", controllerName, deploymentStatus.Message))
		}
	}

	helmReleaseStatus := checkFluxCDHelmReleases(k8sClient)
	if helmReleaseStatus.Status != StatusCompleted {
		overallStatus = helmReleaseStatus.Status
		statusMessages = append(statusMessages, helmReleaseStatus.Message)
	}

	var message string
	if overallStatus == StatusCompleted {
		message = "All FluxCD controllers are running successfully and Helm releases are up to date"
	} else {
		message = strings.Join(statusMessages, "; ")
	}

	return ComponentStatus{
		Name:    "FluxCD",
		Status:  overallStatus,
		Message: fmt.Sprintf("FluxCD: %s", message),
		Step:    1,
	}
}

func checkControllerStatus(k8sClient client.Client, controllerName string) ComponentStatus {
	var deployment appsv1.Deployment
	err := k8sClient.Get(context.Background(), client.ObjectKey{Name: controllerName, Namespace: "flux-system"}, &deployment)
	if err != nil {
		log.Printf("Failed to get deployment %s: %v", controllerName, err)
		return newComponentStatus(controllerName, StatusFailed, fmt.Sprintf("Failed to get deployment: %v", err))
	}

	if deployment.Status.ReadyReplicas == 0 {
		return newComponentStatus(controllerName, StatusInProgress, "Deployment is starting")
	}

	if deployment.Status.ReadyReplicas < *deployment.Spec.Replicas {
		return newComponentStatus(controllerName, StatusInProgress, fmt.Sprintf("Deployment is scaling up (%d/%d ready)", deployment.Status.ReadyReplicas, *deployment.Spec.Replicas))
	}

	return newComponentStatus(controllerName, StatusCompleted, "Deployment is available")
}

func checkFluxCDHelmReleases(k8sClient client.Client) ComponentStatus {
	var helmReleases helmv2.HelmReleaseList
	if err := k8sClient.List(context.Background(), &helmReleases, &client.ListOptions{Namespace: "flux-system"}); err != nil {
		return ComponentStatus{
			Status:  StatusFailed,
			Message: fmt.Sprintf("Failed to list FluxCD Helm releases: %v", err),
		}
	}
	
	if len(helmReleases.Items) == 0 {
		return ComponentStatus{
			Status:  StatusPending,
			Message: "No FluxCD Helm releases found",
		}
	}

	allCompleted := true
	var failedReleases []string

	for _, release := range helmReleases.Items {
		fmt.Println(release.Status.Conditions)
		status, _ := getHelmReleaseStatus(k8sClient, release)
		if status != StatusCompleted {
			allCompleted = false
			if status == StatusFailed {
				failedReleases = append(failedReleases, release.Name)
			}
		}
	}

	if allCompleted {
		return ComponentStatus{
			Status:  StatusCompleted,
			Message: "All FluxCD Helm releases are up to date",
		}
	}

	if len(failedReleases) > 0 {
		return ComponentStatus{
			Status:  StatusFailed,
			Message: fmt.Sprintf("Some FluxCD Helm releases failed: %s", strings.Join(failedReleases, ", ")),
		}
	}

	return ComponentStatus{
		Status:  StatusInProgress,
		Message: "Some FluxCD Helm releases are still being reconciled",
	}
}

func newComponentStatus(name, status, message string) ComponentStatus {
	return ComponentStatus{
		Name:    name,
		Status:  status,
		Message: message,
	}
}

func listHelmReleases(k8sClient client.Client) []helmv2.HelmRelease {
	var helmReleases helmv2.HelmReleaseList
	if err := k8sClient.List(context.Background(), &helmReleases); err != nil {
		log.Printf("Error fetching HelmReleases: %v", err)
		return nil
	}
	return helmReleases.Items
}

func printComponentStatuses(fluxCDStatus ComponentStatus, coreComponentsStatus []ComponentStatus) {
	fmt.Printf("FluxCD Status (Step %d): %s\n", fluxCDStatus.Step, fluxCDStatus.Status)
	fmt.Printf("FluxCD Message: %s\n\n", fluxCDStatus.Message)

	fmt.Println("Core Components Status:")
	for _, component := range coreComponentsStatus {
		fmt.Printf("Step %d: %s (%s): %s\n", component.Step, component.Name, component.Title, component.Status)
		fmt.Printf("Message: %s\n\n", component.Message)
	}
}
