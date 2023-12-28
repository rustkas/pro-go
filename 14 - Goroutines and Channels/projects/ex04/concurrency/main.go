package main

import (
	"fmt"
)

func main() {
	calcTax := func(price float64) float64 {
		return price + (price * 0.2)
	}

	resultChannel := make(chan float64)
	go func (price float64, c chan float64) {
		c <- calcTax(price)
		}(275, resultChannel)
	result := <-resultChannel
	fmt.Println("Result:", result)
}
