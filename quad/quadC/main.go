package main

import (
	"os"
	"piscine"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		return
	}
	piscine.QuadC(x, y)
}
