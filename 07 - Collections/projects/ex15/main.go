package main

import (
	"fmt"
)

func main() {
	names := make([]string, 6, 6)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	names[3] = "Kayak"
	names[4] = "Lifejacket"
	names[5] = "Paddle"
	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))
}
