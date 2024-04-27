package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Investment amount: ")
	// fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Number of years: ")
	// fmt.Print("Number of years: ")
	fmt.Scan(&years)

	outputText("Expected return rate: ")
	// fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFVReal := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)

	// 	fmt.Println("Future Value: ", futureValue)
	// fmt.Printf(`Future Value: %.1f\n
	// Future Value (adjusted for inflation): %.1f\n`, futureValue, futureRealValue)
	// 	fmt.Println("Future value adjsted for inflation", futureRealValue)
	fmt.Print(formattedFV, formattedFVReal)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = investmentAmount * math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
