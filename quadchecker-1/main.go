package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	// 1. Read ALL input from stdin
	inputBytes, _ := io.ReadAll(os.Stdin)
	input := string(inputBytes)

	// If input is empty → not a quad
	if len(input) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// 2. Split input into lines
	lines := strings.Split(input, "\n")

	// Remove last empty line if exists (very common because of newline)
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	// If no valid lines → invalid
	if len(lines) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// 3. Get dimensions
	height := len(lines)
	width := len(lines[0])

	// Check all lines have same width (otherwise invalid shape)
	for _, line := range lines {
		if len(line) != width {
			fmt.Println("Not a quad function")
			return
		}
	}

	// 4. Store matches
	matches := []string{}

	// 5. Compare with each quad
	if quadA(width, height) == input {
		matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", width, height))
	}
	if quadB(width, height) == input {
		matches = append(matches, fmt.Sprintf("[quadB] [%d] [%d]", width, height))
	}
	if quadC(width, height) == input {
		matches = append(matches, fmt.Sprintf("[quadC] [%d] [%d]", width, height))
	}
	if quadD(width, height) == input {
		matches = append(matches, fmt.Sprintf("[quadD] [%d] [%d]", width, height))
	}
	if quadE(width, height) == input {
		matches = append(matches, fmt.Sprintf("[quadE] [%d] [%d]", width, height))
	}

	// 6. Print result
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Sort alphabetically
	sort.Strings(matches)

	// Join with " || "
	fmt.Println(strings.Join(matches, " || "))
}

//
// =======================
// QUAD FUNCTIONS
// =======================
// Each function recreates the drawing
//

func quadA(x, y int) string {
	var result strings.Builder

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {

			// corners
			if (i == 0 || i == y-1) && (j == 0 || j == x-1) {
				result.WriteString("o")
			} else if i == 0 || i == y-1 {
				result.WriteString("-")
			} else if j == 0 || j == x-1 {
				result.WriteString("|")
			} else {
				result.WriteString(" ")
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}

func quadB(x, y int) string {
	var result strings.Builder

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {

			if (i == 0 && j == 0) || (i == y-1 && j == 0) {
				result.WriteString("/")
			} else if (i == 0 && j == x-1) || (i == y-1 && j == x-1) {
				result.WriteString("\\")
			} else if i == 0 || i == y-1 {
				result.WriteString("*")
			} else if j == 0 || j == x-1 {
				result.WriteString("*")
			} else {
				result.WriteString(" ")
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}

func quadC(x, y int) string {
	var result strings.Builder

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {

			if i == 0 && j == 0 {
				result.WriteString("A")
			} else if i == 0 && j == x-1 {
				result.WriteString("A")
			} else if i == y-1 && j == 0 {
				result.WriteString("C")
			} else if i == y-1 && j == x-1 {
				result.WriteString("C")
			} else if i == 0 || i == y-1 || j == 0 || j == x-1 {
				result.WriteString("B")
			} else {
				result.WriteString(" ")
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}

func quadD(x, y int) string {
	var result strings.Builder

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {

			if i == 0 && j == 0 {
				result.WriteString("A")
			} else if i == 0 && j == x-1 {
				result.WriteString("C")
			} else if i == y-1 && j == 0 {
				result.WriteString("A")
			} else if i == y-1 && j == x-1 {
				result.WriteString("C")
			} else if i == 0 || i == y-1 || j == 0 || j == x-1 {
				result.WriteString("B")
			} else {
				result.WriteString(" ")
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}

func quadE(x, y int) string {
	var result strings.Builder

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {

			if i == 0 && j == 0 {
				result.WriteString("A")
			} else if i == 0 && j == x-1 {
				result.WriteString("C")
			} else if i == y-1 && j == 0 {
				result.WriteString("C")
			} else if i == y-1 && j == x-1 {
				result.WriteString("A")
			} else if i == 0 || i == y-1 || j == 0 || j == x-1 {
				result.WriteString("B")
			} else {
				result.WriteString(" ")
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}
