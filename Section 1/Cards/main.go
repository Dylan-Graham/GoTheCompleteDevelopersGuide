package main

func main() {
	cards := newDeck()

	cards.saveToFile("cards.txt")

	cardsFromFile := newDeckFromFile("cardsFromAFile.txt")
	cardsFromFile.print()

}
