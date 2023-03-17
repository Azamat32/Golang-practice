package main

func (b board) IsValid(num, row, col int) bool {
	// check row
	for j := 0; j < boardSize; j++ {
		if b[row][j] == num && j != col {
			return false
		}
	}
	// check column
	for i := 0; i < boardSize; i++ {
		if b[i][col] == num && i != row {
			return false
		}
	}
	// check block
	blockRow := (row / blockSize) * blockSize
	blockCol := (col / blockSize) * blockSize
	for i := blockRow; i < blockRow+blockSize; i++ {
		for j := blockCol; j < blockCol+blockSize; j++ {
			if b[i][j] == num && (i != row || j != col) {
				return false
			}
		}
	}
	return true
}
