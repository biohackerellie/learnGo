package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("There is no file, dumb ass")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Could not parse the fucking file")
	}
	return balance, nil
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		// panic("FUUUUUUUUUUUUUUUUUUUUCK")
	}

	fmt.Print("Welcome to the Bitch Bank!\n")
	for {
		fmt.Println("What do you want to do? bitch")
		fmt.Println("1. Check blance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. withdrawl money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your blance is", accountBalance)
		case 2:
			fmt.Print("How much will you deposit you cheap hoe?: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("wtf, it has to be more than nothing, idiot")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance, "so rich!")
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("How much do you need, bitch: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("You cant withdrawl a negative, moron")
				continue
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Honey, you're too broke for that")
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Ouch, your new balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Bye bitch!")
			fmt.Println("Thanks for choosing the bitch bank, bitch!")
			return
		}
	}
}
