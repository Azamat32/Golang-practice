package piscine

func MakeRange(min, max int) []int {
	arr := []int{}
	if max <= min {
		return nil
	}
	arr = make([]int, max-min)

	if min < max {
		for i := 0; i < (max - min); i++ {
			arr[i] = min + i
		}
	}

	return arr
}
