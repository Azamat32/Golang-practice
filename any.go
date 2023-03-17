package piscine

func Any(f func(string) bool, a []string) bool {
	res1 := true
	res2 := false

	for i := 0; i < len(a); i++ {
		if f(a[i]) == true {
			return res1
		}
	}
	return res2
}
