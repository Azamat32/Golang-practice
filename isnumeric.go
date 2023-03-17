package piscine

func IsNumeric(s string) bool {
	flag := true
	for _, r := range s {
		if r < '0' || r > '9' {
			flag = false
		}
	}
	return flag
}
