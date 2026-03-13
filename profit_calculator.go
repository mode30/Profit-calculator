package main

import (
	"fmt"
)

func main() {
	// var revenue float64
	// var expense float64
	// var taxRate float64

	revenue := getUserInput("enter revenue:")
	expenses := getUserInput("expenses")
	taxRate := getUserInput("tax rate:")
	ebt, profit, ratio := calFinancial(revenue, expenses, taxRate)

	fmt.Println("earning before tax;:", ebt)
	fmt.Println("ratio:", ratio)
	fmt.Println("profit:", profit)
}

func getUserInput(prompt string) float64 {
	var userInput float64
	fmt.Print(prompt)
	fmt.Scan(&userInput)

	return userInput

}

func calFinancial(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	actualTaxRate := taxRate / 100
	profit = ebt * (1 - actualTaxRate)
	ratio = ebt / profit

}
