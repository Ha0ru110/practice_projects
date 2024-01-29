package main

import (
	"fmt"
	"math"
)

func main() {

	var investmentAmount, expectedReturnRate, years, inflationRate float64
	message := ""
	message = welcomeMessage(message)
	fmt.Print(message)
	fmt.Println("How much would you want to invest? (Dollars)")
	fmt.Scan(&investmentAmount)
	fmt.Println("What is the expected annual return rate? (percentage number)")
	fmt.Scan(&expectedReturnRate)
	fmt.Println("Time period? (amount of years)")
	fmt.Scan(&years)
	fmt.Println("Potential inflation rate? (percentage number)")
	fmt.Scan(&inflationRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	formattedFV := fmt.Sprintf("Future Value: %.0f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value considering inflation: %.0f", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func welcomeMessage(a string) string {
	a = "WELCOME TO INVESTMENT CALCULATOR! \nTO GET STARTED, PLEASE ENTER THE PARAMETERS BELOW"
	return a
}
