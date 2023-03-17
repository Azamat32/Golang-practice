package piscine

func Map(f func(int) bool, a []int) []bool {
	arr := []bool{}
	arr = make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		arr[i] = f(a[i])
	}

	return arr
}
