package piscine

func IsPrintable(s string) bool {
	flag := true
	for _, r := range s {
		if r > 127 || r < 32 {
			flag = false
		}
	}
	return flag
}
