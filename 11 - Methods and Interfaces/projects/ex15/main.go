package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	expenses := []Expense {
		Service {"Boat Cover", 12, 89.50, []string{} },
		Service {"Paddle Protect", 12, 8, []string{} },
		}
		for _, expense := range expenses {
		s := expense.(Service)
		fmt.Println("Service:", s.description, "Price:",
		s.monthlyFee * float64(s.durationMonths))
		}
}
