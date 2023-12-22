package main

import (
	"fmt"
)

func main() {
	first := 100
	second := &first
	third := &first
	alpha := 100
	beta := &alpha
	fmt.Println(*second == *third)
	fmt.Println(*second == *beta)
}
