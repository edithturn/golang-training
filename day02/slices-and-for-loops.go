package main

import "fmt"

func main() {

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	fmt.Println("=== Iterating the List ===")
	for i, card := range cards {
		fmt.Println(i, card)
	}

	for value := range cards {
		fmt.Println(value)
	}

	fmt.Println("=== All the List ===")
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}

// Run: go run slices-and-for-loops.go
