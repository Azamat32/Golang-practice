package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	ret := []string{}
	r := []rune(str)
	voz := []rune{}
	for i := 0; i < len(r); i++ {
		if r[i] != ' ' {
			voz = append(voz, r[i])
		} else {
			ret = append(ret, string(voz))
			voz = make([]rune, 0)
		}
		if i == len(r)-1 {
			ret = append(ret, string(voz))
		}
	}
	kek := make(map[string]int)
	for _, num := range ret {
		kek[num] = kek[num] + 1
	}
	if len(r) == 0 {
		return map[string]int{"": 1}
	}
	return kek
}
