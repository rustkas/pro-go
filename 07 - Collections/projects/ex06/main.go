package main

import (
	"fmt"
)

func main() {
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	otherArray := names
	names[0] = "Canoe"
	fmt.Println("names:", names)
	fmt.Println("otherArray:", otherArray)
}
