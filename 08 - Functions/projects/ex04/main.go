package main

import (
	"fmt"
)

func printPrice(product string, price, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

func main() {
	printPrice("Kayak", 275, 0.2)
	printPrice("Lifejacket", 48.95, 0.2)
	printPrice("Soccer Ball", 19.50, 0.15)
}
