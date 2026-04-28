package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		return
	}

	insert := ""
	order := false
	str := ""
	i := 0

	for i < len(args) {
		arg := args[i]

		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		}

		// handle --insert=value
		if len(arg) > 9 && arg[:9] == "--insert=" {
			insert = arg[9:]
			i++
			continue
		}

		// handle -i=value
		if len(arg) > 3 && arg[:3] == "-i=" {
			insert = arg[3:]
			i++
			continue
		}

		// handle --insert value OR -i value
		if arg == "--insert" || arg == "-i" {
			i++
			if i < len(args) {
				insert = args[i]
				i++
				continue
			}
		}

		if arg == "--order" || arg == "-o" {
			order = true
			i++
			continue
		}

		// else it's the main string
		str = arg
		i++
	}

	// process result
	result := str + insert

	if order {
		runes := []rune(result)
		for j := 0; j < len(runes); j++ {
			for k := 0; k < len(runes)-1-j; k++ {
				if runes[k] > runes[k+1] {
					runes[k], runes[k+1] = runes[k+1], runes[k]
				}
			}
		}
		result = string(runes)
	}

	// print result
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printHelp() {
	help := "--insert\n" +
		"  -i\n" +
		"\t This flag inserts the string into the string passed as argument.\n" +
		"--order\n" +
		"  -o\n" +
		"\t This flag will behave like a boolean, if it is called it will order the argument.\n"

	for _, r := range help {
		z01.PrintRune(r)
	}
}
