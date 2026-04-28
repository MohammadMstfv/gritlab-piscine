package piscine

import "github.com/01-edu/z01"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isSafe(positions *[8]int, col, row int) bool {
	for c := 0; c < col; c++ {
		r := positions[c]
		if r == row || abs(r-row) == col-c {
			return false
		}
	}
	return true
}

func solve(col int, positions *[8]int) {
	if col == 8 {
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(positions[i] + '0'))
		}
		z01.PrintRune('\n')
		return
	}
	for row := 1; row <= 8; row++ {
		if isSafe(positions, col, row) {
			positions[col] = row
			solve(col+1, positions)
		}
	}
}

func EightQueens() {
	var positions [8]int
	solve(0, &positions)
}
