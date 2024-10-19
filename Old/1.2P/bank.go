package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

var (
	choice                 int
	err                    error
	accountBalance, amount float64
)

func main() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
	for {
		accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
		if err != nil {
			fmt.Println("ERROR:", err)
			fmt.Println("-------------------")
			// panic("can't continue,  SORRY!")
		}
		presentOptions()
		choice = getUserChoice()
		performOperation(choice)
		if choice == 4 {
			break
		}
	}
	fmt.Println("Thanks for choosing our Bank!")
}

func getUserChoice() int {
	fmt.Print("Enter your Choice: ")
	fmt.Scan(&choice)
	return choice
}

func performOperation(choice int) {
	// if wantCheckBalance := choice == 1; wantCheckBalance
	switch choice {
	case 1:
		fmt.Println("Your Balance is", accountBalance)
	case 2:
		fmt.Print("Amount to deposit: ")
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0")
			return
		}
		accountBalance += amount
		fileops.WriteBalanceToFile(accountBalanceFile, accountBalance)
		fmt.Println("Balance Updated! New Amount is", accountBalance)
	case 3:
		fmt.Print("Amount to Withdraw: ")
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0")
			return
		}
		if amount > accountBalance {
			fmt.Println("Invalid amount. You can't withdraw more than you have.")
			return
		}
		accountBalance -= amount
		fileops.WriteBalanceToFile(accountBalanceFile, accountBalance)
		fmt.Println("Balance Updated! New Amount is", accountBalance)
	default:
		fmt.Println("Goodbye!")
		return
	}
}
