package main

import (
	"fmt"
	"math"
)


func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Number of years: ")
	fmt.Scan(&years)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years) 
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years) 

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFVReal := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)

//	fmt.Println("Future Value: ", futureValue)
	// fmt.Printf()
	// fmt.Println("Future value adjsted for inflation", futureRealValue)
}