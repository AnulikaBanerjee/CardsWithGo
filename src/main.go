package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	hand, remainingDeck := cards.deal(5)
	hand.printCards()
	fmt.Println("--------")
	remainingDeck.printCards()

}
