package main

import (
	"fmt"
	"github.com/profitCalculator/fileOps"
)

func main() {

	var userChoice int
	var accountBalance, err  = fileops.ReadFromFile("balance.txt")

	if(err != nil){
		fmt.Println(err)
	}

	for {

		userOptions()

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

			fileops.WriteBalanceToFile(accountBalance, "balance.txt")
			fmt.Printf("You have withdrawn %v amount, Your total account balance is %v ", customerWithrawalAmount, accountBalance)

		} else if userChoice == 3 {
			var customerDeposit float64
			fmt.Println("Enter Amount to deposit")
			fmt.Scan(&customerDeposit)

			accountBalance += customerDeposit

			fileops.WriteBalanceToFile(accountBalance, "balance.txt")
			fmt.Printf("You have deposited %v, Your current balance is %v", customerDeposit, accountBalance)
		} else {
			break
		}

	}

	fmt.Println("Thank you for banking with us")
}

