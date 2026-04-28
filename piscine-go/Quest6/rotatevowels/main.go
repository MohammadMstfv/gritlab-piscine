package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func main() {
	args := os.Args[1:]

	// If no arguments → just newline
	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Convert all args into rune slices
	allRunes := [][]rune{}
	vowels := []rune{}

	for _, arg := range args {
		r := []rune(arg)
		allRunes = append(allRunes, r)

		for _, ch := range r {
			if isVowel(ch) {
				vowels = append(vowels, ch)
			}
		}
	}

	// Reverse vowels
	for i := 0; i < len(vowels)/2; i++ {
		vowels[i], vowels[len(vowels)-1-i] = vowels[len(vowels)-1-i], vowels[i]
	}

	// Put vowels back
	index := 0
	for i := range allRunes {
		for j := range allRunes[i] {
			if isVowel(allRunes[i][j]) {
				allRunes[i][j] = vowels[index]
				index++
			}
		}
	}

	// Print result with spaces
	for i, word := range allRunes {
		for _, r := range word {
			z01.PrintRune(r)
		}
		if i != len(allRunes)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
