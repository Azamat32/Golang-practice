package piscine

func Unmatch(a []int) int {
	b := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				b++
			}
		}
		if b%2 == 1 {
			return a[i]
		}
	}

	return -1
}
