package main

import (
	"fmt" //means format, similar to Java’s System.out and Scanner
	"hourlypay/utils"
)

func main() {
	pound := "£" //shorthand for declaring and initializing a variable
	var hourlyRate, itemPrice float64
	//this is so strange u can declare 2 variables on one line. float64 is javas double

	fmt.Print("Enter your hourly rate: " + pound + "\n")
	fmt.Scan(&hourlyRate)

	fmt.Print("Enter the item's price:" + pound + "\n")
	fmt.Scan(&itemPrice) //I think & is used to modify the variable here.

	hours := CalculateWorkHours(hourlyRate, itemPrice)
	fmt.Printf("You need to work %.2f hours to afford this item.\n", hours)
	//%.2f formats the float to 2 decimal places
}
