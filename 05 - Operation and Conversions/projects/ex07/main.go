package main

import (
	"fmt"
)

func main() {
	first := 100
	const second = 200.00
	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanOrEqual := first <= second
	greaterThan := first > second
	greaterThanOrEqual := first >= second
	fmt.Println(equal)
	fmt.Println(notEqual)
	fmt.Println(lessThan)
	fmt.Println(lessThanOrEqual)
	fmt.Println(greaterThan)
	fmt.Println(greaterThanOrEqual)
}
