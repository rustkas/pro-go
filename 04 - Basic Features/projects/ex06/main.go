package main

import (
	"fmt"
)

func main() {
	const (
		Watersports = iota
		Soccer
		Chess
	)
	fmt.Println("Watersports:", Watersports)
	fmt.Println("Soccer:", Soccer)
	fmt.Println("Chess:", Chess)
}
