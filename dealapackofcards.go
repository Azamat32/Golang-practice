package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	player1 := "Player 1: "
	player2 := "Player 2: "
	player3 := "Player 3: "
	player4 := "Player 4: "
	printda(player1)
	printlen(0, 3, deck)
	z01.PrintRune('\n')

	printda(player2)
	printlen(3, 6, deck)

	z01.PrintRune('\n')

	printda(player3)
	printlen(6, 9, deck)

	z01.PrintRune('\n')

	printda(player4)
	printlen(9, 12, deck)

	z01.PrintRune('\n')
}

func printda(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
}

func printlen(a, b int, deck []int) {
	for i := a; i < b; i++ {
		fmt.Printf("%v", deck[i])

		if i != b-1 {
			fmt.Printf(", ")
		}

	}
}
