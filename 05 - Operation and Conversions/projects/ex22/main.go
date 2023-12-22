package main

import (
	"fmt"
	"strconv"
)

func main() {
	var val1 string
	val1 = "48.95"
	var float1 float64
	var float1err error
	float1, float1err = strconv.ParseFloat(val1, 64)
	if float1err == nil {
		fmt.Println("Parsed value:", float1)
	} else {
		fmt.Println("Cannot parse", val1, float1err)
	}

	val1 = "4.895e+01"
	float1, float1err = strconv.ParseFloat(val1, 64)
	if float1err == nil {
		fmt.Println("Parsed value:", float1)
	} else {
		fmt.Println("Cannot parse", val1, float1err)
	}
}
