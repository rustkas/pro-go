package main

import (
	"fmt"
)

func main() {
	names := [...]string { "Kayak", "Lifejacket", "Paddle" }
	names[0] = "Hello"
	// var otherArray [4]string
	// names = [3]string(otherArray)
	fmt.Println(names)
}
