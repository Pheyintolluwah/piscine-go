package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		start := i * 3
		fmt.Printf("Player %d: %d, %d, %d\n", i+1, deck[start], deck[start+1], deck[start+2])
	}
}
