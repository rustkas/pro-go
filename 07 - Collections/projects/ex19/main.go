package main

import (
	"fmt"
)

func main() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	allNames := products[:]
	someNames := allNames[:]
	for i := 0; i < 1000; i++ {
		someNames = append(someNames, "Gloves")
	}

	fmt.Println("products", products)
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
}
