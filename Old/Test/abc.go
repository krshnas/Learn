package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	helmv2 "github.com/fluxcd/helm-controller/api/v2beta1"
// 	appsv1 "k8s.io/api/apps/v1"
// 	metav1 "k8s.io/apimachinery/pkg/runtime"
// 	"sigs.k8s.io/controller-runtime/pkg/client"
// 	"sigs.k8s.io/controller-runtime/pkg/client/config"
// 	metav1client "k8s.io/apimachinery/pkg/apis/meta/v1"
// )

// // Model for component status
// type ComponentStatus struct {
// 	Name    string
// 	Title   string
// 	Status  string
// 	Message string
// }

// func main() {
// 	// Set up the Kubernetes client
// 	cfg, err := config.GetConfig()
// 	if err != nil {
// 		log.Fatalf("Error getting in-cluster config: %v", err)
// 	}

// 	scheme := metav1.NewScheme()
// 	err = helmv2.AddToScheme(scheme)
// 	if err != nil {
// 		log.Fatalf("Error adding HelmRelease to scheme: %v", err)
// 	}

// 	// Add apps/v1 scheme for checking Deployment status (for Flux CD)
// 	err = appsv1.AddToScheme(scheme)
// 	if err != nil {
// 		log.Fatalf("Error adding apps/v1 scheme: %v", err)
// 	}

// 	k8sClient, err := client.New(cfg, client.Options{Scheme: scheme})
// 	if err != nil {
// 		log.Fatalf("Error creating Kubernetes client: %v", err)
// 	}

// 	// Fetch components by labels and group them
// 	infraComponents := fetchComponentsByLabel(k8sClient, "infra")
// 	runtimeComponents := fetchComponentsByLabel(k8sClient, "runtime")

// 	// Fetch FluxCD status separately
// 	fluxCDStatus := getFluxCDStatus(k8sClient)

// 	// Print FluxCD status
// 	fmt.Println("FluxCD Status:")
// 	fmt.Printf("Name: %s, Status: %s, Message: %s\n", fluxCDStatus.Name, fluxCDStatus.Status, fluxCDStatus.Message)

// 	// Print infra components
// 	fmt.Println("\nInfrastructure Components:")
// 	for _, component := range infraComponents {
// 		fmt.Printf("Name: %s, Type: %s, Status: %s, Message: %s\n", component.Name, component.Title, component.Status, component.Message)
// 	}

// 	// Print runtime components
// 	fmt.Println("\nRuntime Components:")
// 	for _, component := range runtimeComponents {
// 		fmt.Printf("Name: %s, Type: %s, Status: %s, Message: %s\n", component.Name, component.Title, component.Status, component.Message)
// 	}
// }

// // fetchComponentsByLabel retrieves components by the given component-type (infra or runtime) and returns a list of ComponentStatus
// func fetchComponentsByLabel(k8sClient client.Client, componentType string) []ComponentStatus {
//     var components []ComponentStatus

//     // Create label selector for the given component type
//     labelSelector := &metav1client.LabelSelector{
//         MatchLabels: map[string]string{
//             "component-type": componentType,
//         },
//     }

//     // Convert metav1.LabelSelector to labels.Selector
//     selector, err := metav1client.LabelSelectorAsSelector(labelSelector)
//     if err != nil {
//         log.Printf("Error converting label selector: %v", err)
//         return components
//     }

//     // List all HelmReleases filtered by the label selector
//     var helmReleases helmv2.HelmReleaseList
//     listOptions := &client.ListOptions{
//         LabelSelector: selector,
//     }
//     err = k8sClient.List(context.Background(), &helmReleases, listOptions)
//     if err != nil {
//         log.Printf("Error fetching HelmReleases for %s: %v", componentType, err)
//         return components
//     }

//     // Process each HelmRelease
//     for _, hr := range helmReleases.Items {
//         status := getHelmReleaseStatus(hr)
//         message := getMessageByStatus(status)

//         components = append(components, ComponentStatus{
//             Name:    hr.Name,
//             Title:   componentType,
//             Status:  status,
//             Message: message,
//         })
//     }

//     return components
// }

// // getFluxCDStatus retrieves the status of FluxCD from the deployment status
// func getFluxCDStatus(k8sClient client.Client) ComponentStatus {
// 	var fluxDeployment appsv1.Deployment
// 	err := k8sClient.Get(context.Background(), client.ObjectKey{
// 		Namespace: "flux-system", // Assuming Flux is in the 'flux-system' namespace
// 		Name:      "source-controller", // Example deployment name for FluxCD
// 	}, &fluxDeployment)
// 	if err != nil {
// 		log.Printf("Error fetching Flux CD deployment status: %v", err)
// 		return ComponentStatus{
// 			Name:    "FluxCD",
// 			Status:  "Failed",
// 			Message: fmt.Sprintf("Error: %v", err),
// 		}
// 	}

// 	status := getDeploymentStatus(&fluxDeployment)
// 	message := getMessageByStatus(status)

// 	return ComponentStatus{
// 		Name:    "FluxCD",
// 		Status:  status,
// 		Message: message,
// 	}
// }

// // getDeploymentStatus returns the status of a Deployment
// func getDeploymentStatus(deployment *appsv1.Deployment) string {
// 	if deployment.Status.Replicas == 0 {
// 		return "Pending"
// 	} else if deployment.Status.ReadyReplicas == *deployment.Spec.Replicas {
// 		return "Completed"
// 	} else if deployment.Status.ReadyReplicas < *deployment.Spec.Replicas && deployment.Status.ReadyReplicas > 0 {
// 		return "In Progress"
// 	} else {
// 		return "Failed"
// 	}
// }

// // getHelmReleaseStatus retrieves the status of a HelmRelease by checking the ReleaseCondition
// func getHelmReleaseStatus(hr helmv2.HelmRelease) string {
// 	for _, condition := range hr.Status.Conditions {
// 		if condition.Type == helmv2.ReleasedCondition {
// 			switch condition.Reason {
// 			case "InstallSucceeded":
// 				return "Completed"
// 			case "InstallFailed":
// 				return "Failed"
// 			case "Progressing":
// 				return "In Progress"
// 			}
// 		}
// 	}
// 	return "Pending"
// }

// // getMessageByStatus returns a success or error message based on the status
// func getMessageByStatus(status string) string {
// 	switch status {
// 	case "Completed":
// 		return "Installation succeeded"
// 	case "Failed":
// 		return "Installation failed"
// 	case "In Progress":
// 		return "Installation in progress"
// 	default:
// 		return "Installation pending"
// 	}
// }