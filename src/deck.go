package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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

func (d deck) shuffle() {
	// random seed creation code added later --
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var temp int
	for i := range d {
		//temp = rand.Intn(len(d)-1) This will produce the same shuffle result
		//because golang uses the same seed evertime. So we need something where we can provide a random seed.
		temp = r.Intn(len(d) - 1)
		d[i], d[temp] = d[temp], d[i]

	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) printCards() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1) //os.Exit ends the program. Exit 0 means normal ending and any other number is error.
	}

	return deck(strings.Split(string(bs), (",")))
}
