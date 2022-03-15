package main

import "fmt"

func main() {
	card := newCard()
	cardInt := newCardInt()

	fmt.Println(card)
	fmt.Println(cardInt)

}

func newCard() string {
	return "Five of Diamonds"
}

func newCardInt() int {
	return 100
}
