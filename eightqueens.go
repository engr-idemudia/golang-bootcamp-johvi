package piscine

import (
	"github.com/01-edu/z01"
)

func printTable(table []int) {
	for _, v := range table {
		z01.PrintRune(rune(48 + v))
	}
	z01.PrintRune('\n')
}

func canPlace(table []int, row, col int) bool {
	// Check if position not under fire
	for i := 0; i < 8; i++ {
		rowsBefore := row - i
		// Check if any in same column
		if table[i] == col {
			return false
		}
		if rowsBefore > 0 {
			// Check diagonals
			if table[row-rowsBefore] == col-rowsBefore || table[row-rowsBefore] == col+rowsBefore {
				return false
			}
		}

	}
	return true
}

func solveQueens(table []int, row int) {
	if row == 8 {
		printTable(table)
		return
	}
	// For each column
	for col := 1; col <= 8; col++ {
		if canPlace(table, row, col) {
			table[row] = col
			solveQueens(table, row+1)
		}
		table[row] = 0
	}
}

func EightQueens() {
	solveQueens([]int{0, 0, 0, 0, 0, 0, 0, 0}, 0)
}
