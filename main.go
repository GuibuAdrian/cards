package main

func main() {
	cards := newDeck()
	/*
		// hand, remainingDeck := <deck>, <deck>
		hand, remainingCards := deal(cards, 5)

		hand.print()
		remainingCards.print()

		fmt.Println(cards.toString())

		cards.saveToFile("my_cards")

		cards = newDeckFromFile("my_cards")

		cards.print()
	*/

	cards.shuffle()
	cards.print()
}
