package main

import (
	"fmt" //means format, similar to Java’s System.out and Scanner
	"hourly-pay-item-calculator/utils/calculator"
)

func main() {
	const pound = "£"
	var hourlyRate, itemPrice float64

	fmt.Print("Enter your hourly rate: " + pound + "\n")
	fmt.Scan(&hourlyRate)

	fmt.Print("Enter the item's price:" + pound + "\n")
	fmt.Scan(&itemPrice) // & is used to modify the variable here.

	hours := calculator.BasicCalculator{}.Calculate(hourlyRate, itemPrice)
	fmt.Printf("You need to work %.2f hours to afford this item.\n", hours)
}
