package main

import (
	"bufio"
	"fmt"
	"os"
	"piscine/quads"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input strings.Builder
	for scanner.Scan() {
		input.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input")
		os.Exit(1)
	}
	inp := input.String()
	if len(inp) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	lines := strings.Split(strings.TrimSuffix(inp, "\n"), "\n")
	height := len(lines)
	if height == 0 {
		fmt.Println("Not a quad function")
		return
	}
	width := len(lines[0])
	for _, line := range lines {
		if len(line) != width {
			fmt.Println("Not a quad function")
			return
		}
	}
	if width == 0 || height == 0 {
		fmt.Println("Not a quad function")
		return
	}
	var matches []string
	quadFuncs := []struct {
		name string
		fn   func(int, int) string
	}{
		{"quadA", quads.QuadA},
		{"quadB", quads.QuadB},
		{"quadC", quads.QuadC},
		{"quadD", quads.QuadD},
		{"quadE", quads.QuadE},
	}
	for _, q := range quadFuncs {
		if q.fn(width, height) == inp {
			matches = append(matches, fmt.Sprintf("[quad%s] [%d] [%d]", strings.ToUpper(q.name[4:]), width, height))
		}
	}
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		sort.Strings(matches)
		fmt.Println(strings.Join(matches, " || "))
	}
}
