package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solve(board, 0)
}

func solve(board [8]int, col int) {
	if col == 8 {
		printSolution(board)
		return
	}

	for row := 0; row < 8; row++ {
		if isSafe(board, col, row) {
			board[col] = row
			solve(board, col+1)
		}
	}
}

func isSafe(board [8]int, col int, row int) bool {
	for prevCol := 0; prevCol < col; prevCol++ {
		prevRow := board[prevCol]
		if prevRow == row {
			return false
		}
		if prevRow-prevCol == row-col || prevRow+prevCol == row+col {
			return false
		}
	}
	return true
}

func printSolution(board [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + 1 + '0'))
	}
	z01.PrintRune('\n')
}
