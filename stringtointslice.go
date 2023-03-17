package piscine

func StringToIntSlice(str string) []int {
	arr := []int{}
	if len(str) == 0 {
		arr = nil
	}
	for _, s := range str {
		arr = append(arr, int(rune(s)))
	}

	return arr
}
