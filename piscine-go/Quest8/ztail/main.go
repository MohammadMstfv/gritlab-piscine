package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 || args[0] != "-c" {
		return
	}

	count := 0
	for _, c := range args[1] {
		if c >= '0' && c <= '9' {
			count = count*10 + int(c-'0')
		}
	}

	files := args[2:]
	hasError := false

	for i, name := range files {
		data, err := os.ReadFile(name)
		if err != nil {
			fmt.Printf("%v\n", err)
			hasError = true
			continue
		}

		// Print header ONLY if multiple files AND file is valid
		if len(files) > 1 {
			if i > 0 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", name)
		}

		start := 0
		if len(data) > count {
			start = len(data) - count
		}

		fmt.Printf("%s", data[start:])
	}

	if hasError {
		os.Exit(1)
	}
}
