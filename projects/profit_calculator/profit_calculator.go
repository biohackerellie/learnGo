package main

import (
	"fmt"
	"os"
)

func main() {

	revenue := printScan("Revenue: ")

	expenses := printScan("Expenses: ")

	taxRate := printScan("Tax rate: ")

	EBT, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	printResults(EBT, profit, ratio)

}

func writeToFile(results string) {
	os.WriteFile("results.txt", []byte(results), 0644)
}

func printScan(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		fmt.Println("Please enter a number greater than 0")
		panic("FUUUUUUUUUUUUUUUUUUUUUUUCK")
	}
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

	results := fmt.Sprintf("Your EBT is: %.1f\nYour profit is: %.3f\nYour EBT to profit ratio is: %.1f\n", EBT, profit, ratio)
	writeToFile(results)
}
