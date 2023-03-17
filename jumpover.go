package piscine

func JumpOver(str string) string {
	newStr := ""
	for i := 2; i < len(str); i += 3 {
		if len(str) > 2 {
			newStr += string(str[i])
		}
	}

	return newStr + "\n"
}
