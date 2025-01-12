package main

import (
	"fmt"
	currencyFmt "packages/fmt"
	"packages/store"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", currencyFmt.ToCurrency(product.Price()))
}

//
