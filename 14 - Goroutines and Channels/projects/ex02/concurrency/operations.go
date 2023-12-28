package main

import (
	"fmt"
	"time"
)

// func CalcStoreTotal(data ProductData) {
// 	var storeTotal float64
// 	for category, group := range data {
// 		storeTotal += group.TotalPrice(category)
// 	}
// 	fmt.Println("Total:", ToCurrency(storeTotal))
// }

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64)
	for category, group := range data {
		go group.TotalPrice(category, channel)
	}
	for i := 0; i < len(data); i++ {
		storeTotal += <-channel
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

// func (group ProductGroup) TotalPrice(category string) (total float64) {

// 	for _, p := range group {

// 		fmt.Println("-->", category, "product:", p.Name)
// 		total += p.Price
// 		time.Sleep(time.Millisecond * 500)
// 	}
// 	fmt.Println(category, "subtotal:", ToCurrency(total))
// 	return
// }

func (group ProductGroup) TotalPrice(category string, resultChannel chan float64) {
	var total float64
	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println(category, "subtotal:", ToCurrency(total))
	resultChannel <- total
}
