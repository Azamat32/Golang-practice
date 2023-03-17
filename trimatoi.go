package piscine

func TrimAtoi(s string) int {
	a := []rune(s)
	result := 0
	decimal := 1
	c := 0
	flag := false
	for _, word := range a {
		if word >= '0' && word <= '9' {
			for i := '0'; i < word; i++ {
				c++
			}
			result = result*10 + c

			c = 0
		}
		if word == '-' && result == 0 {
			flag = true
		}
	}
	if flag == true {
		decimal = -1
	}
	return result * decimal
}
