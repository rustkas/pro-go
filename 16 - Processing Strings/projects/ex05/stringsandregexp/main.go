package main

// import (
// 	"fmt"
// 	// "strings"
// 	"golang.org/x/text/cases"
// 	"golang.org/x/text/language"
// )

// func main() {
// 	description := "A boat for sailing"
// 	fmt.Println("Original:", description)

// 	// Создаем объект TitleCase для английского языка
// 	tc := cases.Title(language.English)
// 	fmt.Println("Title:", tc.String(description))
// 	// fmt.Println("Title:", strings.Title(description))
// }

import (
	"fmt"
	"strings"
)

func main() {
	specialChar := "\u01c9"
	fmt.Println("Original:", specialChar, []byte(specialChar))
	upperChar := strings.ToUpper(specialChar)
	fmt.Println("Upper:", upperChar, []byte(upperChar))
	titleChar := strings.ToTitle(specialChar)
	fmt.Println("Title:", titleChar, []byte(titleChar))
}