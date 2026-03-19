package main

import "fmt"

func main() {
	var productName string // Declare a variable but no value is assigned
	var productPrice int=1000 // Declare a variable and assign value
	var companyName = "Apple Inc." // Declare a variable and assign a value but type is not declared

	// *** mostly used ***
	category:="Electronics" // Declare a variable and assign a value but type is not declared and var is not used

	// declare a const variable
	const availableStock=50

	fmt.Println("Product Name is ",productName," and price is ",productPrice," and company name is ",companyName, " and category is ",category," and Stock available is ",availableStock)
}
