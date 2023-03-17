package piscine

func AppendRange(min, max int) []int {
	arr := []int{}
	if max <= min {
		return nil
	}
	if min < max {
		for i := min; i < max; i++ {
			arr = append(arr, i)
		}
	}

	return arr
}
