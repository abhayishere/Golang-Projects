package main

// import "fmt"

func main() {
	// cards := newDeck()
	// hand, remaingCards := deal(cards, 5)
	// hand.print()
	// remaingCards.print()
	// cards:=newDeck()
	// cards.saveToFile("my_cards")
	cards:=newDeckFromFile("my_cards")
	cards.print()
	cards.shuffle()
	cards.print()
}
