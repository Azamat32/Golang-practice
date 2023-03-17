package piscine

func IsAlphaT(s string) bool {
	flag := true
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			flag = false
		}
	}
	return flag
}

func Capitalize(s string) string {
	s = ToLowerT(s)
	for i, v := range s {
		if i == 0 {
			s = ToUpperT(string(v)) + s[i+1:]
		} else {
			if IsAlphaT(string(v)) && !IsAlphaT(string(s[i-1])) {
				if i != len(s)-1 {
					s = s[:i] + ToUpperT(string(v)) + s[i+1:]
				} else {
					s = s[:i] + ToUpperT(string(v))
				}
			}
		}
	}
	return s
}

func ToLowerT(s string) string {
	var result string
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			result += string(c + 32)
		} else {
			result += string(c)
		}
	}
	return result
}

func ToUpperT(s string) string {
	var result string
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			result += string(c - 32)
		} else {
			result += string(c)
		}
	}
	return result
}
