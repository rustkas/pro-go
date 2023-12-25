package main

import (
	"fmt"
)

func main() {
	names := []string {"Kayak", "Lifejacket", "Paddle"}
	names = append(names, "Hat", "Gloves")
	fmt.Println(names)
}
