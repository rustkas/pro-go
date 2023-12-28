package main

import (
	"fmt"
	"packages/store"
)

func main() {
	// product := store.Product{
	// 	Name:     "Kayak",
	// 	Category: "Watersports",
	// 	// price: 279,
	// }
	product := store.NewProduct("Kayak", "Watersports", 279)

	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", product.Price())
}

//
