package main

import (
	"fmt"
	"os"
)

const revenue = "revenue.txt"

func fileSaver(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", EBT, profit, ratio)
	os.WriteFile(revenue, []byte(results), 0644)
}

var EBT, profit, ratio float64

func main() {

	EBT, profit, ratio = calculations(EBT, profit, ratio)

	fileSaver(EBT, profit, ratio)

	fileSaver(EBT, profit, ratio)
	fmt.Print("EBT:  ")
	fmt.Printf("%.1f\n", EBT)
	fmt.Print("Profit:  ")
	fmt.Printf("%.1f\n", profit)
	fmt.Print("Ratio:  ")
	fmt.Printf("%.1f", ratio)
}

func calculations(revenue, expenses, taxRate float64) (EBT, profit, ratio float64) {
	revenue, expenses, taxRate = getAllParameters(revenue, expenses, taxRate)
	EBT = revenue - expenses
	profit = EBT * (1 - taxRate/100)
	ratio = EBT / profit
	return EBT, profit, ratio
}

func getAllParameters(revenue, expenses, taxRate float64) (revenue2, expenses2, taxRate2 float64) {
	fmt.Print("WELCOME TO REVENUE CALCULATOR", "\nEnter revenue:")
	fmt.Scan(&revenue)
	if revenue <= 0 {
		fmt.Println("Negative numbers or 0 not accepted")
		panic("System exited")
	}
	revenue2 = revenue
	fmt.Print("Enter expenses:")
	fmt.Scan(&expenses)
	if expenses < 0 {
		fmt.Println("Negative numbers not accepted")
		panic("System exited")
	}
	expenses2 = expenses
	fmt.Print("Enter tax rate:")
	fmt.Scan(&taxRate)
	if taxRate < 0 || taxRate > 100 {
		fmt.Println("Negative numbers or numbers over 100 not accepted")
		panic("System exited")
	}
	taxRate2 = taxRate
	return revenue2, expenses2, taxRate2
}
