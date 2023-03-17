package piscine

func Rot14(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'l' || s[i] >= 'A' && s[i] <= 'L' {
			result += string(s[i] + 14)
		} else if s[i]+14 > 'z' {
			result += string(s[i] - 26 + 14)
		} else if s[i]+14 > 'Z' {
			result += string(s[i] - 26 + 14)
		} else {
			result += string(s[i])
		}
	}

	return result
}
