package piscine

func BasicAtoi(s string) int {
	a := []rune(s)
	result := 0

	for _, r := range a {
		n := int(r - '0')
		result = result*10 + n
	}
	return result
}
