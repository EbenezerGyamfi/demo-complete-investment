package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var userChoice int
	var accountBalance, err  = readBalanceFromFile()

	if(err != nil){
		panic("server error, contact admin")
	}

	for {

		fmt.Println("Welcome to Go bank Console")
		fmt.Println("Please select what you want")
		fmt.Println("1. Account")
		fmt.Println("2. Withdrawal")
		fmt.Println("3. Deposit")
		fmt.Println("4. Exit")

		fmt.Scan(&userChoice)

		if userChoice == 1 {
			fmt.Println("You entered 1")
			fmt.Printf("Your account Balance is %v", accountBalance)
		} else if userChoice == 2 {
			var customerWithrawalAmount float64
			fmt.Println("Enter withdrawal Amount ")
			fmt.Scan(&customerWithrawalAmount)

			if customerWithrawalAmount <= 0 {
				fmt.Println("Invalid input, try again")
				continue
			}

			if customerWithrawalAmount > accountBalance {
				fmt.Println("Insufficient Amount, please try again")
				continue
			}

			accountBalance -= customerWithrawalAmount

			writeBalanceToFile(accountBalance)
			fmt.Printf("You have withdrawn %v amount, Your total account balance is %v ", customerWithrawalAmount, accountBalance)

		} else if userChoice == 3 {
			var customerDeposit float64
			fmt.Println("Enter Amount to deposit")
			fmt.Scan(&customerDeposit)

			accountBalance += customerDeposit

			writeBalanceToFile(accountBalance)
			fmt.Printf("You have deposited %v, Your current balance is %v", customerDeposit, accountBalance)
		} else {
			break
		}

	}

	fmt.Println("Thank you for banking with us")
}

func writeBalanceToFile(balance float64) {

	balanceText := fmt.Sprint(balance)

	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func readBalanceFromFile() (float64, error) {

	data, err := os.ReadFile("balance.txt")


	if(err != nil){
		return 1000,  errors.New("no such file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if(err != nil){
		return 1000, errors.New("no such file")
	}
	return balance, nil
}
