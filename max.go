package piscine

func Max(a []int) int {
	c := 0

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a)-1; j++ {
			if a[i] < a[j] {
				c = a[j]
			}
		}
	}

	return c
}
