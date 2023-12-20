package main

import (
	"fmt"
)

func main() {
	price, tax, inStock, discount := 275.00, 27.50, true,
		true
	var salesPerson = "Alice"
	fmt.Println("Total:", price+tax)
	fmt.Println("In stock:", inStock)
}
