package main

import (
	"os"

	"github.com/01-edu/z01"
)

func stringToInt(s string) (int, bool) {
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}
		n = n*10 + int(r-'0')
	}
	return n, true
}

func main() {
	args := os.Args[1:]
	upper := false

	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	if len(args) == 0 {
		return
	}

	for _, arg := range args {
		n, valid := stringToInt(arg)
		if !valid || n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}

		letter := rune('a' + n - 1)
		if upper {
			letter = rune('A' + n - 1)
		}
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
