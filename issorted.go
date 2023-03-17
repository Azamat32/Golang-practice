package piscine

func f(a, b int) int {
	res := 1
	if a < b {
		res = -1
		return res
	} else if a == b {
		res = 0
		return res

	}

	return res
}

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) > 1 {
		if f(a[0], a[1]) <= 0 {
			for i := 0; i < len(a)-1; i++ {
				if f(a[i], a[i+1]) > 0 {
					return false
				}
			}
		}

		if f(a[0], a[1]) >= 0 {
			for i := 0; i < len(a)-1; i++ {
				if f(a[i], a[i+1]) < 0 {
					return false
				}
			}
		}
	}

	return true
}
