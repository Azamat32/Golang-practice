package piscine

func Atoi(s string) int {
	a := []rune(s)
	result := 0
	decimal := 1
	flag := false
	for i := 0; i < len(a); i++ {

		n := int(a[i] - '0')
		if n >= 0 && n <= 9 {
			result = result*10 + n
		} else if a[i] == '-' || a[i] == '+' {
			if !flag {
				flag = true
				if a[i] == '-' {
					decimal = -1
				} else {
					decimal = 1
				}
			} else {
				return 0
			}
		} else {
			return 0
		}
	}
	if flag && (a[len(a)-1] == '-' || a[len(a)-1] == '+') {
		return 0
	}
	return result * decimal
}
