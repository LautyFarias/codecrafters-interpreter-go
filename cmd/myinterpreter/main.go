package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func printErrorAndExit(error string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, error, args)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		printErrorAndExit("Usage: ./your_program.sh tokenize <filename>")
	}

	command := os.Args[1]

	if command != "tokenize" {
		printErrorAndExit("Unknown command: %s\n", command)
	}

	filename := os.Args[2]
	file, err := os.Open(filename)

	defer file.Close()

	if err != nil {
		printErrorAndExit("Error opening file: %v\n", err)
	}

	reader := bufio.NewReader(file)
	bytes := make([]byte, 1)

	for {
		_, err := reader.Read(bytes)

		if err == io.EOF {
			fmt.Println("EOF  null")
			break
		}

		panic("Scanner not implemented")
	}
}
