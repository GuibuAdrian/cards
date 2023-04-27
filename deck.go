package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{"Ace of Spades", "Two of Spades", newCard()}
	return cards
}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
