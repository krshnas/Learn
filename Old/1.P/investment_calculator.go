package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investAmount, expectedReturnRate, years float64 // default value of float64 will be assigned here, 0.0
	// years := 10.0	// another way to initialize variable

	// fmt.Print("Investemt Amount: ")
	outputText("Investemt Amount: ")
	fmt.Scan(&investAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)

	// Outputs Information
	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Value (adjusted for Inflation): ", futureRealValue)
	// fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for Inflation): %.1f", futureValue, futureRealValue)
	// fmt.Printf(`Future Value: %.1f
	// Future Value (adjusted for Inflation): %.1f`, futureValue, futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}
