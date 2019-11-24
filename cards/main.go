package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards.txt")
	myDeck := newDeckFromFile("my_cards")
	myDeck.shuffle()
	myDeck.print()
}
