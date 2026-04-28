package main

import (
	"fmt"
	"os"
)

// Constants for sudoku dimensions
const (
	GRID_SIZE = 9  // Size of the sudoku grid (9x9)
	BOX_SIZE  = 3  // Size of each box (3x3)
	EMPTY     = 0  // Represents empty cell
	GIVEN     = -1 // Marker for cells that are given in the input
)

// Matrix represents a sudoku grid
type Matrix [GRID_SIZE][GRID_SIZE]int

// Candidates tracks possible values for each cell
type Candidates [GRID_SIZE][GRID_SIZE][]int

// parseInput converts string arguments to a sudoku matrix
// Each argument represents a row, dots represent empty cells
func parseInput(args []string) (*Matrix, error) {
	// Check if we have exactly 9 arguments (9 rows)
	if len(args) != GRID_SIZE {
		return nil, fmt.Errorf("Error")
	}

	matrix := &Matrix{}

	// Parse each row
	for i := 0; i < GRID_SIZE; i++ {
		if len(args[i]) != GRID_SIZE {
			return nil, fmt.Errorf("Error")
		}

		// Parse each cell in the row
		for j := 0; j < GRID_SIZE; j++ {
			char := args[i][j]

			// Check if it's a dot (empty cell)
			if char == '.' {
				matrix[i][j] = EMPTY
			} else if char >= '1' && char <= '9' {
				// Convert character to digit
				digit := int(char - '0')
				matrix[i][j] = digit
			} else {
				// Invalid character
				return nil, fmt.Errorf("Error")
			}
		}
	}

	return matrix, nil
}

// isValidInput checks if the initial sudoku setup is valid (no duplicate numbers in rows, columns, or boxes)
func isValidInput(matrix *Matrix) bool {
	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			if matrix[i][j] != EMPTY {
				value := matrix[i][j]
				// Temporarily remove the value
				matrix[i][j] = EMPTY

				// Check if this value can be placed here
				if !isSafe(matrix, i, j, value) {
					return false
				}

				// Restore the value
				matrix[i][j] = value
			}
		}
	}
	return true
}

// isSafe checks if placing a number at position (row, col) is safe, Returns true if the number doesn't conflict with existing numbers
func isSafe(matrix *Matrix, row, col, num int) bool {
	return isSafeRow(matrix, row, num) &&
		isSafeCol(matrix, col, num) &&
		isSafeBox(matrix, row, col, num)
}

// isSafeRow checks if the number already exists in the row
func isSafeRow(matrix *Matrix, row, num int) bool {
	for col := 0; col < GRID_SIZE; col++ {
		if matrix[row][col] == num {
			return false
		}
	}
	return true
}

// isSafeCol checks if the number already exists in the column
func isSafeCol(matrix *Matrix, col, num int) bool {
	for row := 0; row < GRID_SIZE; row++ {
		if matrix[row][col] == num {
			return false
		}
	}
	return true
}

// isSafeBox checks if the number already exists in the 3x3 box
func isSafeBox(matrix *Matrix, row, col, num int) bool {
	// Calculate the top-left corner of the 3x3 box
	boxRow := (row / BOX_SIZE) * BOX_SIZE
	boxCol := (col / BOX_SIZE) * BOX_SIZE

	// Check all cells in the 3x3 box
	for i := boxRow; i < boxRow+BOX_SIZE; i++ {
		for j := boxCol; j < boxCol+BOX_SIZE; j++ {
			if matrix[i][j] == num {
				return false
			}
		}
	}
	return true
}

// buildCandidates initializes candidate lists for all empty cells, Uses constraint propagation to eliminate impossible values
func buildCandidates(matrix *Matrix) *Candidates {
	candidates := &Candidates{}

	// Initialize candidates for each cell
	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			if matrix[i][j] == EMPTY {
				// Build list of possible values for this empty cell
				for num := 1; num <= 9; num++ {
					if isSafe(matrix, i, j, num) {
						candidates[i][j] = append(candidates[i][j], num)
					}
				}
			}
		}
	}
	return candidates
}

// removeCandidateValue removes a value from candidates in affected cells
// Constraint propagation: when a number is placed, remove it from peers
func removeCandidateValue(candidates *Candidates, row, col, num int) {
	// Remove from row
	for j := 0; j < GRID_SIZE; j++ {
		removeFromList(&candidates[row][j], num)
	}

	// Remove from column
	for i := 0; i < GRID_SIZE; i++ {
		removeFromList(&candidates[i][col], num)
	}

	// Remove from box
	boxRow := (row / BOX_SIZE) * BOX_SIZE
	boxCol := (col / BOX_SIZE) * BOX_SIZE
	for i := boxRow; i < boxRow+BOX_SIZE; i++ {
		for j := boxCol; j < boxCol+BOX_SIZE; j++ {
			removeFromList(&candidates[i][j], num)
		}
	}
}

// removeFromList removes a value from a candidate list
func removeFromList(list *[]int, value int) {
	for i, v := range *list {
		if v == value {
			*list = append((*list)[:i], (*list)[i+1:]...)
			return
		}
	}
}

// getNumberPriority returns numbers sorted by frequency in given cells
// Numbers that appear more are tried first (fewer slots left to fill)
func getNumberPriority(matrix *Matrix) []int {
	count := make(map[int]int)

	// Count occurrences of each given number
	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			if matrix[i][j] != EMPTY {
				count[matrix[i][j]]++
			}
		}
	}

	// Create list of numbers sorted by count (descending)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if count[numbers[i]] < count[numbers[j]] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	return numbers
}

// solveWithOptimization applies optimization techniques before backtracking
func solveWithOptimization(matrix *Matrix) bool {
	// Build candidate lists for constraint propagation
	candidates := buildCandidates(matrix)

	// Get number priority based on frequency in given cells
	numberPriority := getNumberPriority(matrix)

	// Use backtracking with candidate optimization
	return solveBacktrackOptimized(matrix, candidates, numberPriority)
}

// solveBacktrackOptimized uses backtracking with MRV heuristic
// MRV (Minimum Remaining Values): choose cell with fewest candidates first
// Returns true if a solution is found, false otherwise
func solveBacktrackOptimized(matrix *Matrix, candidates *Candidates, numberPriority []int) bool {
	// Find the empty cell with minimum remaining values (MRV heuristic)
	minCandidates := 10
	minRow, minCol := -1, -1

	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			// Found an empty cell
			if matrix[i][j] == EMPTY {
				// Check if this cell has fewer candidates
				if len(candidates[i][j]) < minCandidates {
					minCandidates = len(candidates[i][j])
					minRow, minCol = i, j
				}

				// If no candidates, puzzle is unsolvable
				if minCandidates == 0 {
					return false
				}
			}
		}
	}

	// No empty cells found - puzzle solved!
	if minRow == -1 {
		return true
	}

	// Try each candidate value for this cell
	for _, num := range candidates[minRow][minCol] {
		// Place the number
		matrix[minRow][minCol] = num

		// Save current candidates state for backtracking
		savedCandidates := [GRID_SIZE][GRID_SIZE][]int{}
		for i := 0; i < GRID_SIZE; i++ {
			for j := 0; j < GRID_SIZE; j++ {
				savedCandidates[i][j] = append([]int{}, candidates[i][j]...)
			}
		}

		// Remove this number from candidates of affected cells (constraint propagation)
		removeCandidateValue(candidates, minRow, minCol, num)

		// Recursively try to solve
		if solveBacktrackOptimized(matrix, candidates, numberPriority) {
			return true
		}

		// Backtrack: restore matrix and candidates
		matrix[minRow][minCol] = EMPTY
		for i := 0; i < GRID_SIZE; i++ {
			for j := 0; j < GRID_SIZE; j++ {
				candidates[i][j] = append([]int{}, savedCandidates[i][j]...)
			}
		}
	}

	// No valid number found, backtrack
	return false
}

/*
// solveBacktrack uses backtracking algorithm to solve the sudoku
// Returns true if a solution is found, false otherwise
func solveBacktrack(matrix *Matrix) bool {
	// Find the next empty cell
	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			// Found an empty cell
			if matrix[i][j] == EMPTY {
				// Try numbers 1-9
				for num := 1; num <= 9; num++ {
					if isSafe(matrix, i, j, num) {
						// Place the number
						matrix[i][j] = num

						// Recursively try to solve
						if solveBacktrack(matrix) {
							return true
						}

						// Backtrack if no solution found
						matrix[i][j] = EMPTY
					}
				}
				// No valid number found, backtrack
				return false
			}
		}
	}
	// No empty cells left - puzzle solved!
	return true
}
*/

// printMatrix displays the solved sudoku
func printMatrix(matrix *Matrix) {
	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}

// main function - entry point of the program
func main() {
	// Check if we have the correct number of arguments
	if len(os.Args) != GRID_SIZE+1 {
		fmt.Println("Error")
		return
	}

	// Parse the input into a matrix
	matrix, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Println("Error")
		return
	}

	// Validate the initial input
	if !isValidInput(matrix) {
		fmt.Println("Error")
		return
	}

	// Solve the sudoku with optimization techniques
	if solveWithOptimization(matrix) {
		// Successfully solved
		printMatrix(matrix)
	} else {
		// No solution found
		fmt.Println("Error")
	}
}
