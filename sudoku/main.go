package main

import (
	"os"

	"github.com/01-edu/z01"
)

const (
	boardSize = 9
	blockSize = 3
)

type board [boardSize][boardSize]int

func main() {
	if len(os.Args) != boardSize+1 {
		z01.PrintRune('E')
		z01.PrintRune('r')
		z01.PrintRune('r')
		z01.PrintRune('o')
		z01.PrintRune('r')
		z01.PrintRune('$')

		z01.PrintRune('\n')
		return
	}
	// check if each row has boardSize cells
	for i := 1; i < len(os.Args); i++ {
		if len(os.Args[i]) != boardSize {
			z01.PrintRune('E')
			z01.PrintRune('r')
			z01.PrintRune('r')
			z01.PrintRune('o')
			z01.PrintRune('r')
			z01.PrintRune('$')

			z01.PrintRune('\n')
			return
		}
	}
	input := os.Args[1:]
	var b board
	for i, row := range input {
		for j, ch := range row {
			if ch != '.' {
				b[i][j] = int(ch - '0')
			}
		}
	}
	solution, ok := solve(b, 0, 0)
	if ok {
		solution.Print()
	} else {
		z01.PrintRune('E')
		z01.PrintRune('r')
		z01.PrintRune('r')
		z01.PrintRune('o')
		z01.PrintRune('r')
		z01.PrintRune('$')

		z01.PrintRune('\n')
	}
}
