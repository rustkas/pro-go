package main

import (
	"fmt"
)

func main() {
	first := 100
	var second *int
	// fmt.Println(*second)
	second = &first
	second = nil
	fmt.Println(*second)
	fmt.Println(second == nil)
}
