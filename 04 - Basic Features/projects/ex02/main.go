package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Go")
	fmt.Println(20 + 20)
	fmt.Println(20 + 30)
	fmt.Println(20.2, -20.2, 1.2e10, 1.2e-10)
	fmt.Println(true, false)
	fmt.Println('A', '\n', '\u00A5', '¥')
}
