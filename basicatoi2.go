package piscine

func BasicAtoi2(s string) int {
	a := []rune(s)
	result := 0

	for _, r := range a {

		n := int(r - '0')
		if n >= 0 && n <= 9 {
			result = result*10 + n
		} else {
			return 0
		}
	}
	return result
}
