package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards.shuffle()

	hand, remainingDeck := cards.deal(5)
	hand.printCards()
	fmt.Println("--------")
	//remainingDeck.printCards()
	remainingDeck.saveToFile("remaining")
	newDeck := readDeckFromFile("remaining")
	hand2, _ := newDeck.deal(4)
	hand2.printCards()

}
