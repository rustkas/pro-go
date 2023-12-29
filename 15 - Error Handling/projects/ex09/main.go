package main

import "fmt"

type CategoryCountMessage struct {
	Category string
	Count    int
}

func processCategories(categories []string, outChan chan<- CategoryCountMessage) {
	defer func() {
		if arg := recover(); arg != nil {
			fmt.Println(arg)
			close(outChan)
		}
	}()
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			outChan <- CategoryCountMessage{
				Category: message.Category,
				Count:    int(message.Total),
			}
		} else {
			panic(message.CategoryError)
		}
	}
	close(outChan)
}

func main() {
	categories := []string{"Watersports", "Chess", "Running"}
	channel := make(chan CategoryCountMessage)
	go processCategories(categories, channel)
	for message := range channel {
		fmt.Println(message.Category, "Total:", message.Count)
	}
}
