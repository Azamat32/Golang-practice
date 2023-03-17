package piscine

func FirstRune(s string) rune {
	var first rune
	for _, c := range s {
		first = c
		break
	}
	return first
}
