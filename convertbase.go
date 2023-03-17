package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	res := AtoiBase(nbr, baseFrom)
	str1 = ""
	PrintNbrBase2(res, baseTo)
	return str1
}

var str1 string = ""

func PrintNbrBase2(nbr int, base string) {
	ParseBase(nbr, []rune(base), len(base), 1)
}

func ParseBase(nbr int, base []rune, size int, dot int) {
	if size > nbr && nbr > size*-1 {
		str1 += string(base[nbr*dot])
		return
	}
	ParseBase(nbr/size, base, size, dot)
	str1 += string(base[nbr%size*dot])
}

func AtoiBase(nbr string, base string) int {
	size := 0
	isValid := true
	for _, val := range base {
		if val == '-' || val == '+' {
			isValid = false
			size = 0
			break
		}
		size++
	}
	if size < 2 {
		return 0
	}
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if base[j] == base[i] {
				isValid = false
				break
			}
		}
	}
	if !isValid {
		return 0
	}
	for _, val := range nbr {
		cnt := 0
		for _, val1 := range base {
			if val == val1 {
				cnt++
			}
		}
		if cnt == 0 {
			return 0
		}
	}
	return ParseToInt([]rune(nbr), []rune(base), len(base))
}

func GetIndex(ch rune, base []rune) int {
	for i := 0; i < len(base); i++ {
		if base[i] == ch {
			return i
		}
	}
	return 0
}

func ParseToInt(nbr []rune, base []rune, size int) int {
	ans := 0
	for i := 0; i < len(nbr); i++ {
		ans = ans + (IterativePower(len(base), len(nbr)-1-i))*GetIndex(nbr[i], base)
	}
	return ans
}
