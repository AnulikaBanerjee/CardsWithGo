package main

import (
	"fmt"
	"strconv"
)

type deck []string

func newDeck() deck {
	types := []string{"Club", "Spades", "Diamonds", "Hearts"}
	fullDeck := deck{}
	for _, typ := range types {
		for i := 1; i <= 13; i++ {
			switch i {
			case 1:
				fullDeck = append(fullDeck, "Ace of "+typ)
			case 11:
				fullDeck = append(fullDeck, "Jack of "+typ)
			case 12:
				fullDeck = append(fullDeck, "Queen of"+typ)
			case 13:
				fullDeck = append(fullDeck, "King of "+typ)
			default:
				fullDeck = append(fullDeck, (strconv.Itoa(i) + " of " + typ))
			}
		}
	}
	return fullDeck
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) printCards() {
	for _, card := range d {
		fmt.Println(card)
	}
}
