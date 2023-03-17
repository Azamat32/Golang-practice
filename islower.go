package piscine

func IsLower(s string) bool {
	for _, charNumber := range s {
		if charNumber > 122 || charNumber < 97 {
			return false
		}
	}
	return true
}
