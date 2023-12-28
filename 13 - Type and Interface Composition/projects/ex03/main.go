package main

import (
	"composition/store"
	"fmt"
)

func main() {
	// rentals := []*store.RentalBoat{
	// 	store.NewRentalBoat("Rubber Ring", 10, 1, false,
	// 		false, "N/A", "N/A"),
	// 	store.NewRentalBoat("Yacht", 50000, 5, true, true,
	// 		"Bob", "Alice"),
	// 	store.NewRentalBoat("Super Yacht", 100000, 15, true,
	// 		true,
	// 		"Dora", "Charlie"),
	// }
	// for _, r := range rentals {
	// 	fmt.Println("Rental Boat:", r.Name, "Rental Price:",
	// 		r.Price(0.2),
	// 		"Captain:", r.Captain)
	// }

	product := store.NewProduct("Kayak", "Watersports", 279)
	deal := store.NewSpecialDeal("Weekend Special", product, 50)
	Name, price, Price := deal.GetDetails()
	fmt.Println("Name:", Name)
	fmt.Println("Price field:", price)
	fmt.Println("Price method:", Price)
}
