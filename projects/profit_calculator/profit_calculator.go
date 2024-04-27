package main

import (
	"fmt"
)

func main() {

	revenue := printScan("Revenue: ")
	expenses := printScan("Expenses: ")
	taxRate := printScan("Tax rate: ")

	EBT, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	printResults(EBT, profit, ratio)
}

func printScan(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	return userInput
}

func calculateProfit(revenue, expenses, taxRate float64) (EBT, profit, ratio float64) {
	EBT = revenue - expenses
	profit = EBT * (1 - taxRate/100)
	ratio = EBT / profit
	return EBT, profit, ratio
}

func printResults(EBT, profit, ratio float64) {
	fmt.Print("Your EBT is: ")
	fmt.Printf("%.1f\n", EBT)

	fmt.Print("Your profit is: ")
	fmt.Printf("%.3f\n", profit)

	fmt.Print("Your EBT to profit ratio is: ")
	fmt.Printf("%.1f\n", ratio)
}
