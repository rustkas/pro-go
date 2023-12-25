package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
		// otherNames []string
	}
	p1 := Product{name: "Kayak", category: "Watersports",
		price: 275.00}
	p2 := Product{name: "Kayak", category: "Watersports",
		price: 275.00}
	p3 := Product{name: "Kayak", category: "Watersports", price: 275.01}
	fmt.Println("p1 == p2:", p1 == p2)
	fmt.Println("p1 == p3:", p1 == p3)
}
