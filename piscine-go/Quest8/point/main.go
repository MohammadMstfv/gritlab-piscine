package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	z := '0'
	a := z
	a++ // '1'
	b := a
	b++ // '2'
	d := b
	d++
	d++ // '4'

	output := []rune{'x', ' ', '=', ' ', d, b, ',', ' ', 'y', ' ', '=', ' ', b, a, '\n'}

	for _, r := range output {
		z01.PrintRune(r)
	}
}
