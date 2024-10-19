package main

import (
	"errors"
	"fmt"
	"os"
)

const errMsg = "ERROR:"

// const panicMsg = "Can't Continue, SORRY!"

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(errMsg, err)
		// panic(panicMsg)
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(errMsg, err)
		// panic(panicMsg)
		return
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(errMsg, err)
		// panic(panicMsg)
		return
	}
	// if err1 != nil || err2 != nil || err3 != nil{
	// 	fmt.Println(errMsg, err1)
	// 	// panic(panicMsg)
	// 	return
	// }

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	printUserOutput("EBT: %.1f\n", ebt)
	printUserOutput("Profit: %.1f\n", profit)
	printUserOutput("Ratio: %.1f\n", ratio)
}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	status := validateUserInput(userInput)
	if !status {
		return userInput, errors.New("value must be positive number")
	}
	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	storeResults(ebt, profit, ratio)
	return ebt, profit, ratio
}

func printUserOutput(text string, value float64) {
	fmt.Printf(text, value)
}

func validateUserInput(input float64) bool {
	return input > 0
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio:%.3f\n", ebt, profit, ratio)
	os.WriteFile("profit.txt", []byte(fmt.Sprint(results)), 0644)
}
