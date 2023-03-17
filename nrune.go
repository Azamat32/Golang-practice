package piscine

func NRune(s string, n int) rune {
	var a rune
	if n > len(s) || n < 1 {
		return 0
	}
	runes := []rune(s)
	a = runes[n-1]
	return a
}
