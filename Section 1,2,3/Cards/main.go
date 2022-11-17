package main

import "fmt"

// $: go run main.go <user written dependencies>.go

func main() {
	cards := newDeck()

	cards.saveToFile("cards.txt")

	cardsFromFile := newDeckFromFile("cardsFromAFile.txt")
	cardsFromFile.print()

	fmt.Println("Shuffled cards from file:")
	cardsFromFile.shuffle().print()

}
