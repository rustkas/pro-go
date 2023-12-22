package main

import (
	"fmt"
)

func main() {
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime value:", val, counter)
		default:
			fmt.Println("Non-prime value:", val, counter)
		}
	}
}
