package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	var e1 Expense = &Product { name: "Kayak" }
	var e2 Expense = &Product { name: "Kayak" }
	var e3 Expense = Service { description: "Boat Cover" }
	var e4 Expense = Service { description: "Boat Cover" }
	fmt.Println("e1 == e2", e1 == e2)
	fmt.Println("e3 == e4", e3 == e4)
}
