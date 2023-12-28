package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	product := Product{"Kayak", "Watersports", 275}
	var expense Expense = &product
	product.price = 100
	fmt.Println("Product field value:", product.price)
	fmt.Println("Expense method result:", expense.getCost(false))
}
