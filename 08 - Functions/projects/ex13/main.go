package main

import (
	"fmt"
)

func calcTax(price float64) float64 {
	return price + (price * 0.2)
}
func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		fmt.Println("Product: ", product, "Price:", calcTax(price))
	}
}
