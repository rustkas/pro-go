package main

import (
	"fmt"
)

func main() {
	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChannel)
	for {
		// details := <-dispatchChannel
		// fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
		if details, open := <-dispatchChannel; open {
			fmt.Println("Dispatch to", details.Customer, ":",
				details.Quantity,
				"x", details.Product.Name)
		} else {
			fmt.Println("Channel has been closed")
			break
		}
	}
}
