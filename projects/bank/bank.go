package main

import (
	"fmt"

	"epklabs.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		// panic("FUUUUUUUUUUUUUUUUUUUUCK")
	}

	fmt.Print("Welcome to the Bitch Bank!\n")
	fmt.Println(randomdata.PhoneNumber())
	for {
		presentOptions()
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Bye bitch!")
			fmt.Println("Thanks for choosing the bitch bank, bitch!")
			return
		}
	}
}
