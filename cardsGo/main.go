package main

import (
	"fmt"
)

var ACE_OF_SPADES string = "Ace of spades"

var deckSize int

func main() {

	deckSize = 52

	card := "Ace of spades"
	card = "Five of Diamonds"
	cards := []string{newCard(), "5 Diamonds"}
	cards = append(cards, "10 King")
	for i, item := range cards {
		fmt.Println(i, item)
	}
	count := 5
	for i := 0; i < count; i++ {
		fmt.Println("Loop: ", i+1)
	}
	fmt.Println(card)
	fmt.Println(ACE_OF_SPADES)
	fmt.Println(deckSize)
	fmt.Println(newCard())
	fmt.Println(cards)
}

func newCard() string {
	return "NEW CARD"
}

func primitive() {
	s := "this is string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}

func constVar() {
	const myConst = "a"
	fmt.Printf("%v, %T\n", myConst, myConst)
}
