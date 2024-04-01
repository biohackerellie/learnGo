package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	
	fmt.Print("Enter your total revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your total expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your tax rate: ")
	fmt.Scan(&taxRate)


	EBT := revenue - expenses
	profit := EBT * (1 - taxRate / 100)
	ratio := EBT / profit

	fmt.Print("Your EBT is: ")
	fmt.Println(EBT)

	fmt.Print("Your profit is: ")
	fmt.Println(profit)

	fmt.Print("Your EBT to profit ratio is: ")
	fmt.Println(ratio)
}