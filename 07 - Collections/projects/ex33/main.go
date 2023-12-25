package main

import (
	"fmt"
)

func main() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	value, ok := products["Hat"]
	if ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}
}