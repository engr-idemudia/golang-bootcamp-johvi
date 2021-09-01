package main

import (
	"fmt"
	"os"
)

// check if runes in string are numbers or '.'
func ValidArg(str string) bool {
	for i := range str {
		if !((str[i] >= '1' && str[i] <= '9') || str[i] == '.') {
			return false
		}
	}
	return true
}

// check if string is long enough (9), change runes int, check for '.'
func ParseArgs(args []string, result *[9][9]int) bool {
	for i := 1; i < len(args); i++ {
		if len(args[i]) != 9 {
			return false
		} else {
			if ValidArg(args[i]) {
				for j := 0; j < 9; j++ {
					if args[i][j] != '.' {
						result[i-1][j] = int(args[i][j] - 48)
					}
				}
			} else {
				return false
			}
		}
	}
	return true
}

// check for empty cells, start filling with canditate 'cand' numbers from 1 to 9 (the backtracking)
func SolveSud(result *[9][9]int) bool {
	if !HasEmptyCell(*result) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if result[i][j] == 0 {
				for cand := 1; cand <= 9; cand++ {
					result[i][j] = cand
					if IsValidSud(*result) {
						if SolveSud(result) {
							return true
						}
						result[i][j] = 0
					} else {
						result[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

// check if cell value is empty, used in other funcs
func HasEmptyCell(result [9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if result[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

// check for duplicates, to use later
func HasDuplicates(counter [10]int) bool {
	for i := 1; i < 10; i++ {
		if counter[i] > 1 {
			return true
		}
	}
	return false
}

// check duplicates by rows and columns
func IsValidSud(result [9][9]int) bool {
	for i := 0; i < 9; i++ {
		counter := [10]int{}
		counter2 := [10]int{}
		for j := 0; j < 9; j++ {
			counter[result[i][j]]++
			counter2[result[j][i]]++
		}
		if HasDuplicates(counter) {
			return false
		}
		if HasDuplicates(counter2) {
			return false
		}
	}

	// check 3x3 section for duplicates
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[result[row][col]]++
				}
				if HasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

// run the conditions and solve sudoku
func main() {
	args := os.Args
	lenArgs := len(args)
	if lenArgs == 10 {
		var result [9][9]int
		if ParseArgs(args, &result) {
			if SolveSud(&result) {

				for i := 0; i < 9; i++ {
					for j := 0; j < 9; j++ {
						fmt.Print(result[i][j])
						if j != 8 {
							fmt.Print(" ")
						}
					}
					fmt.Println()
				}
				return

			}
		}
	}
	fmt.Println("Error")
}
