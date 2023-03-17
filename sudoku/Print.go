package main

import "github.com/01-edu/z01"

func (b board) Print() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			num := b[i][j]
			if num == 0 {
				z01.PrintRune('.')
			} else {
				z01.PrintRune(rune(num + '0'))
			}
			if j < boardSize-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
