package main

import (
	"fmt"
)

func main() {
	type Product struct {
		name, category string
		price          float64
	}
	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
	}
	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)

	var lifejacket Product
	fmt.Println("Name is zero value:", lifejacket.name == "")
	fmt.Println("Category is zero value:", lifejacket.category == "")
	fmt.Println("Price is zero value:", lifejacket.price == 0)
}
