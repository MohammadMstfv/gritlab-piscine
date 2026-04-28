package main

import "github.com/01-edu/z01"

// var alphabet string = "abcdefghijklmnopqrstuvwxyz"
func main() {
	i := 122
	for i >= 97 {
		z01.PrintRune(rune(i))
		i--
	}
	z01.PrintRune('\n')
}
