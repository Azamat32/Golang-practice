package piscine

func bubblesort(array []int) {
	size := 5
	for step := 0; step < size-1; step++ {
		for i := 0; i < size-step-1; i++ {
			if array[i] > array[i+1] {

				temp := array[i]
				array[i] = array[i+1]
				array[i+1] = temp
			}
		}
	}
}

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	bubblesort(arr)

	return arr[2]
}
