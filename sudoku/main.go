package main

import (
	"fmt"
	"os"
)

func main() {

	// ------------------------ Error Handling ------------------------

	// Check if we have the correct number of arguments (rows)
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	var grid [9][9]int
	for i := 0; i < 9; i++ {
		// counts columns, checks if they are exactly 9
		if len(os.Args[i+1]) != 9 {
			fmt.Println("Error")
			return
		} // ---------- Filling Grid with input and empty cells -------
		for j, c := range os.Args[i+1] {
			if c == '.' {
				grid[i][j] = 0 // empty cells represented as 0
			} else if c >= '1' && c <= '9' {
				grid[i][j] = int(c - '0') // other cells converted to integers

			} else { // checks valid input, error if not valid
				fmt.Println("Error")
				return
			}
		}
	}

	// Validate initial grid
	if !isInitialGridValid(&grid) {
		fmt.Println("Error")
		return
	}

	// Check for Multiple Solutions
	gridCopy := grid
	solutionCount := 0
	countSolutions(&gridCopy, &solutionCount)
	if solutionCount != 1 {
		fmt.Println("Error")
		return
	}

	if !solve(&grid) { // if not solvable, error
		fmt.Println("Error")
		return
	}

	// ------------------------ Print Grid ----------------------------

	for i := 0; i < 9; i++ { // print the visual grid in the terminal
		for j := 0; j < 9; j++ {
			if j > 0 {
				fmt.Print(" ") // print spaces in between every cell printed
			}
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
	// ----------------------------------------------------------------
}

// Check the Input
func isInitialGridValid(grid *[9][9]int) bool {
	// Check rows
	for i := 0; i < 9; i++ {
		seen := make(map[int]bool)
		for j := 0; j < 9; j++ {
			if grid[i][j] != 0 {
				if seen[grid[i][j]] {
					return false // duplicate found
				}
				seen[grid[i][j]] = true
			}
		}
	}

	// Check columns
	for j := 0; j < 9; j++ {
		seen := make(map[int]bool)
		for i := 0; i < 9; i++ {
			if grid[i][j] != 0 {
				if seen[grid[i][j]] {
					return false // duplicate found
				}
				seen[grid[i][j]] = true
			}
		}
	}

	// Check 3x3 boxes
	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			seen := make(map[int]bool)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					row := boxRow*3 + i
					col := boxCol*3 + j
					if grid[row][col] != 0 {
						if seen[grid[row][col]] {
							return false // duplicate found
						}
						seen[grid[row][col]] = true
					}
				}
			}
		}
	}

	return true
}

func solve(grid *[9][9]int) bool {
	for i := 0; i < 9; i++ { // runs for each row
		for j := 0; j < 9; j++ { // runs for each column
			if grid[i][j] == 0 {
				for num := 1; num <= 9; num++ { // checks filling numbers one by one
					if isValid(grid, i, j, num) {
						grid[i][j] = num
						if solve(grid) { // check if it's done completely
							return true
						}
						grid[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(grid *[9][9]int, row, col, num int) bool {
	for x := 0; x < 9; x++ {
		if grid[row][x] == num || grid[x][col] == num {
			return false
		}
	}
	startRow, startCol := row/3*3, col/3*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

// for checking out the multiple solutions , basically the same as solve but instead of returning true when we find a solution, we count the number of solutions and return if it's more than 1. also since the solve is recursive, we need to pass the count as a pointer so that it can be updated across recursive calls.
func countSolutions(grid *[9][9]int, count *int) {
	// Early exit if we already found more than 1 solution
	if *count > 1 {
		return
	}

	for i := 0; i < 9; i++ { // runs for each row
		for j := 0; j < 9; j++ { // runs for each column
			if grid[i][j] == 0 {
				for num := 1; num <= 9; num++ { // checks filling numbers one by one
					if isValid(grid, i, j, num) {
						grid[i][j] = num
						countSolutions(grid, count)
						grid[i][j] = 0
					}
				}
				return // backtrack
			}
		}
	}
	// No empty cells left - found a complete solution
	*count++
}
