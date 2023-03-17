package piscine

func StrLen(s string) int {
	l := 0
	a := []rune(s)
	for i := range a {
		l = i + 1
	}
	return l
}
