package piscine

import "github.com/01-edu/z01"

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i, r1 := range base {
		if r1 == '+' || r1 == '-' {
			return false
		}
		for j, r2 := range base {
			if i != j && r1 == r2 {
				return false
			}
		}
	}
	return true
}

func printNV() {
	z01.PrintRune('N')
	z01.PrintRune('V')
}

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		printNV()
		return
	}
	negative := nbr < 0
	if negative {
		z01.PrintRune('-')
	}
	var num uint64
	if negative {
		num = uint64(-int64(nbr))
	} else {
		num = uint64(nbr)
	}
	if num == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}
	digits := []rune{}
	baseLen := uint64(len(base))
	for num > 0 {
		digit := num % baseLen
		digits = append([]rune{rune(base[digit])}, digits...)
		num /= baseLen
	}
	for _, d := range digits {
		z01.PrintRune(d)
	}
}
