package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type Supplier struct {
	name, city string
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}
func (product *Product) printDetails() {
	fmt.Println("Name:", product.name, "Category:",
		product.category,
		"Price", product.calcTax(0.2, 100))
}
func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}
func (supplier *Supplier) printDetails() {
	fmt.Println("Supplier:", supplier.name, "City:",
		supplier.city)
}

func (supplier *Supplier) printDetails1(showName bool) {
	if (showName) {
	fmt.Println("Supplier:", supplier.name, "City:",
	supplier.city)
	} else {
	fmt.Println("Supplier:", supplier.name)
	}
	}

func main() {
	products := []*Product{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Lifejacket", "Watersports", 48.95),
		newProduct("Soccer Ball", "Soccer", 19.50),
	}
	for _, p := range products {
		p.printDetails()
	}
	suppliers := []*Supplier{
		{"Acme Co", "New York City"},
		{"BoatCo", "Chicago"},
	}
	for _, s := range suppliers {
		s.printDetails()
	}
}
