package piscine

import "github.com/01-edu/z01"

func PrintNbr(mb int) {
	t := 1
	if mb < 0 {
		t = -1
		z01.PrintRune('-')
	}
	if mb != 0 {
		f := (mb / 10) * t
		if f != 0 {
			PrintNbr(f)
		}
		k := ((mb % 10) * t) + '0'
		z01.PrintRune(rune(k))
	} else {
		z01.PrintRune('0')
	}
}
