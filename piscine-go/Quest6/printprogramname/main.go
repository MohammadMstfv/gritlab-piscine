package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	base := name
	for i, r := range name {
		if r == '/' {
			base = name[i+1:]
		}
	}
	for _, r := range base {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
