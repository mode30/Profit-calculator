package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expense float64
	var taxRate float64

	fmt.Print("enter revenue:")
	fmt.Scan(&revenue)

	fmt.Print("expenses:")
	fmt.Scan(&expense)

	fmt.Print("tax rate:")
	fmt.Scan(&taxRate)

	ebt := revenue - expense
	fmt.Println("earning before tax;:", ebt)

	actualTaxRate := taxRate / 100
	profit := ebt * (1 - actualTaxRate)
	fmt.Println("Profit:", ebt)

	ratio := ebt / profit
	fmt.Println("ratio:", ratio)
}
