package piscine

func Sqrt(nb int) int {
	res := 1
	for i := 0; i <= nb; i++ {
		res = i * i
		if res == nb {
			return i
		}

	}
	if res != nb {
		return 0
	}
	return res
}
