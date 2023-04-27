package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// We never use the words this or self by convention we usually refer to the receiver with a single letter
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
