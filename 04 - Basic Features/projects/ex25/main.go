package main

import (
	"fmt"
)

func main() {
	first := 100
	second := &first
	fmt.Println("First:", first)
	first++
	fmt.Println("First:", first)
	*second++
	var myNewPointer *int
	myNewPointer = second
	*myNewPointer++
	fmt.Println("First:", first)
	fmt.Println("Second:", *second)
}
