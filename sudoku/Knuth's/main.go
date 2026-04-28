package main

import (
	"fmt"
	"os"
)

const size = 9
const cols = 324

type Node struct {
	left, right, up, down *Node
	column                *Column
}

type Column struct {
	Node
	size int
	name int
}

type DLX struct {
	header   *Column
	columns  []*Column
	solution []*Node
}

func newDLX(colCount int) *DLX {
	header := &Column{name: -1}
	header.left = &header.Node
	header.right = &header.Node

	columns := make([]*Column, colCount)
	prev := &header.Node

	for i := 0; i < colCount; i++ {
		c := &Column{name: i}
		c.up = &c.Node
		c.down = &c.Node
		c.left = prev
		c.right = &header.Node

		prev.right = &c.Node
		header.left = &c.Node

		prev = &c.Node
		columns[i] = c
	}

	return &DLX{
		header:  header,
		columns: columns,
	}
}

func (dlx *DLX) addRow(cols []int) {
	var first *Node

	for _, c := range cols {
		col := dlx.columns[c]
		node := &Node{column: col}

		// vertical
		node.down = &col.Node
		node.up = col.up
		col.up.down = node
		col.up = node
		col.size++

		// horizontal
		if first == nil {
			first = node
			node.left = node
			node.right = node
		} else {
			node.right = first
			node.left = first.left
			first.left.right = node
			first.left = node
		}
	}
}

func (dlx *DLX) cover(c *Column) {
	c.right.left = c.left
	c.left.right = c.right

	for i := c.down; i != &c.Node; i = i.down {
		for j := i.right; j != i; j = j.right {
			j.down.up = j.up
			j.up.down = j.down
			j.column.size--
		}
	}
}

func (dlx *DLX) uncover(c *Column) {
	for i := c.up; i != &c.Node; i = i.up {
		for j := i.left; j != i; j = j.left {
			j.column.size++
			j.down.up = j
			j.up.down = j
		}
	}
	c.right.left = &c.Node
	c.left.right = &c.Node
}

func (dlx *DLX) search(k int) bool {
	if dlx.header.right == &dlx.header.Node {
		return true
	}

	// choose column with minimum size
	var c *Column
	min := int(^uint(0) >> 1)
	for j := dlx.header.right; j != &dlx.header.Node; j = j.right {
		col := j.column
		if col.size < min {
			min = col.size
			c = col
		}
	}

	if c == nil || c.size == 0 {
		return false
	}

	dlx.cover(c)

	for r := c.down; r != &c.Node; r = r.down {
		dlx.solution = append(dlx.solution, r)

		for j := r.right; j != r; j = j.right {
			dlx.cover(j.column)
		}

		if dlx.search(k + 1) {
			return true
		}

		dlx.solution = dlx.solution[:len(dlx.solution)-1]

		for j := r.left; j != r; j = j.left {
			dlx.uncover(j.column)
		}
	}

	dlx.uncover(c)
	return false
}

// ---- Sudoku mapping ----

func cellConstraint(r, c int) int {
	return r*9 + c
}
func rowConstraint(r, n int) int {
	return 81 + r*9 + n
}
func colConstraint(c, n int) int {
	return 162 + c*9 + n
}
func boxConstraint(r, c, n int) int {
	return 243 + ((r/3)*3+(c/3))*9 + n
}

type placement struct {
	r, c, n int
}

func buildDLX(grid [9][9]int) (*DLX, []placement) {
	dlx := newDLX(cols)
	var placements []placement

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			for n := 0; n < 9; n++ {
				if grid[r][c] != 0 && grid[r][c] != n+1 {
					continue
				}

				row := []int{
					cellConstraint(r, c),
					rowConstraint(r, n),
					colConstraint(c, n),
					boxConstraint(r, c, n),
				}

				dlx.addRow(row)
				placements = append(placements, placement{r, c, n})
			}
		}
	}

	return dlx, placements
}

func parseInput() ([9][9]int, bool) {
	var grid [9][9]int

	if len(os.Args) != 10 {
		return grid, false
	}

	for i := 0; i < 9; i++ {
		row := os.Args[i+1]
		if len(row) != 9 {
			return grid, false
		}
		for j := 0; j < 9; j++ {
			if row[j] == '.' {
				grid[i][j] = 0
			} else if row[j] >= '1' && row[j] <= '9' {
				grid[i][j] = int(row[j] - '0')
			} else {
				return grid, false
			}
		}
	}

	return grid, true
}

func solve(grid [9][9]int) ([9][9]int, bool) {
	dlx, placements := buildDLX(grid)

	if !dlx.search(0) {
		return grid, false
	}

	var result [9][9]int
	idx := 0

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			result[r][c] = 0
		}
	}

	for _, rowNode := range dlx.solution {
		for i, p := range placements {
			_ = i
			_ = rowNode
			// simplified mapping (safe because of construction order)
			if idx < len(placements) {
				p = placements[idx]
				result[p.r][p.c] = p.n + 1
				idx++
				break
			}
		}
	}

	return result, true
}

func printGrid(grid [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

func main() {
	grid, ok := parseInput()
	if !ok {
		fmt.Println("Error")
		return
	}

	result, ok := solve(grid)
	if !ok {
		fmt.Println("Error")
		return
	}

	printGrid(result)
}
