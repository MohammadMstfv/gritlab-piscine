package main

import (
	"bufio"
	"io"
	"os"
)

func printError(err error) {
	os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
	os.Exit(1)
}

func readAndPrint(r io.Reader) {
	reader := bufio.NewReader(r)
	for {
		rune, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			printError(err)
		}
		os.Stdout.WriteString(string(rune))
	}
}

func main() {
	args := os.Args[1:]

	// No arguments → read from stdin
	if len(args) == 0 {
		readAndPrint(os.Stdin)
		return
	}

	// Loop through files
	for _, fileName := range args {
		file, err := os.Open(fileName)
		if err != nil {
			printError(err)
		}
		readAndPrint(file)
		file.Close()
	}
}
