package main

import (
	"github.com/01-edu/z01"
)

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

type point struct {
	x int
	y int
}

func printInt(n int) {
	if n >= 10 {
		printInt(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')

	z01.PrintRune('=')
	z01.PrintRune(' ')

	printInt(points.x)
	z01.PrintRune(',')

	z01.PrintRune(' ')

	z01.PrintRune('y')
	z01.PrintRune(' ')

	z01.PrintRune('=')
	z01.PrintRune(' ')

	printInt(points.y)
	z01.PrintRune('\n')
}
