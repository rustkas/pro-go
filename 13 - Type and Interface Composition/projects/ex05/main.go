package main

import (
	"fmt"
	"composition/store"
	)
	func main() {
	products := map[string]*store.Product {
	"Kayak": store.NewBoat("Kayak", 279, 1, false),
	"Ball": store.NewProduct("Soccer Ball", "Soccer",
	19.50),
	}
	for _, p := range products {
	fmt.Println("Name:", p.Name, "Category:", p.Category,
	"Price:", p.Price(0.2))
	}
	}