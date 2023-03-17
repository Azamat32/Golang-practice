package piscine

func LastRune(s string) rune {
	var last rune
	arr := []rune(s)
	last = arr[len(s)-1]

	return last
}
