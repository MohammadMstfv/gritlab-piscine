package main

import (
	"fmt"
	"os"
)

func printStr(s string) {
	for _, r := range s {
		fmt.Print(string(r))
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printStr("File name missing\n")
	} else if len(args) > 1 {
		printStr("Too many arguments\n")
	} else {
		filename := args[0]
		file, err := os.Open(filename)
		if err != nil {
			printStr("ERROR\n")
			return
		}
		defer file.Close()
		buffer := make([]byte, 1024)
		for {
			n, err := file.Read(buffer)
			if n > 0 {
				for i := 0; i < n; i++ {
					fmt.Print(string(rune(buffer[i])))
				}
			}
			if err != nil {
				break
			}
		}
	}
}
