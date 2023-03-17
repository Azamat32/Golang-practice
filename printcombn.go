package piscine

import "github.com/01-edu/z01"

func show(n int, table [9]int, max [9]int) {
	i := 0
	for i < n {
		z01.PrintRune(rune(table[i]))
	}
	if table[0] != max[0] {
		z01.PrintRune(',')
		z01.PrintRune(' ')

	}
}

func Combination1() {
	table := [9]int{0}
	max := [9]int{}
	for table[0] <= max[0] {
		show(1, table, max)
		table[0]++
	}
}

func PrintCombN(n int) {
	table := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	max := [9]int{}
	if n == 1 {
		Combination1()
	} else {
		i := n - 1
		j := 9
		for i >= 0 {
			max[i] = j
			i--
			j--
		}
		i = n - 1
		for table[0] < max[0] {
			if table[i] != max[i] {
				show(n, table, max)
				table[i]++
			}
			if table[i] == max[i] {
				if table[i-1] != max[i-1] {
					show(n, table, max)
					table[i-1]++
					j = i
					for j < n {
						table[j] = table[j-1] + 1
						j++
					}
					i = n - 1
				}
			}

		}
	}
}
