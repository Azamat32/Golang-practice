package main

func solve(b board, row, col int) (board, bool) {
	if col == boardSize {
		row++
		col = 0
	}

	if row == boardSize {
		return b, true
	}

	if b[row][col] != 0 {
		return solve(b, row, col+1)
	}

	numSolutions := 0

	var lastSolution board

	for num := 1; num <= boardSize; num++ {
		if b.IsValid(num, row, col) {
			b[row][col] = num
			if result, ok := solve(b, row, col+1); ok {
				numSolutions++
				if numSolutions == 1 {
					lastSolution = result
				}
			}
			b[row][col] = 0
		}
		if numSolutions > 1 {
			return b, false
		}
	}
	if numSolutions == 1 {
		return lastSolution, true
	}
	return b, false
}
