package piscine

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				SwapString(&slice[j], &slice[j+1])
			}
		}
	}

	return slice
}

func SwapString1(w1 *string, w2 *string) {
	temp := *w1
	*w1 = *w2
	*w2 = temp
}
