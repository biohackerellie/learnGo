package main

import (
	"fmt"
)

func main() {

	var accountBalance = 1000.0

	fmt.Print("Welcome to the Bitch Bank!\n")
	fmt.Println("What do you want to do? bitch")
	fmt.Println("1. Check blance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. withdrawl money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your blance is", accountBalance)
	} else if choice == 2 {
		fmt.Print("How much will you deposit you cheap hoe?: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Balance updated! New amount: ", accountBalance, "so rich!")
	} else if choice == 3 {
		fmt.Print("How much do you need, bitch: ")
		var withdrawlAmount float64
		fmt.Scan(&withdrawlAmount)
		accountBalance -= withdrawlAmount
		fmt.Println("Ouch, your new balance is: ", accountBalance)
	}

	fmt.Println("your choice: ", choice)
}
