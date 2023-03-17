package piscine

func Compact(ptr *[]string) int {
	a := []string{}
	a = make([]string, 0)
	for _, str := range *ptr {
		if str != "" {
			a = append(a, str)
		}
	}

	*ptr = a
	return len(a)
}
