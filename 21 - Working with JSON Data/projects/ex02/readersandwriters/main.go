package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	numbers := [3]int{10, 20, 30}
	var byteArray [5]byte
	copy(byteArray[0:], []byte(names[0]))
	byteSlice := []byte(names[0])
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(names)
	encoder.Encode(numbers)
	encoder.Encode(byteArray)
	encoder.Encode(byteSlice)
	fmt.Print(writer.String())
}
