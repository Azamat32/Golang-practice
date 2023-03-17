package piscine

func DescendAppendRange(max, min int) []int {
	arr := []int{}
	if max <= min {
		return arr
	}
	if min < max {
		for i := max; i > min; i-- {
			arr = append(arr, i)
		}
	}

	return arr
}
