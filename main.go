package main

import "fmt"

func main() {
	cards := newDeck()

	// hand, remainingDeck := <deck>, <deck>
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	fmt.Println(cards.toString())
}
