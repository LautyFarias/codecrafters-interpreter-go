package main

import (
	"fmt"
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
	fileContents, err := os.ReadFile(filename)

	if err != nil {
		printErrorAndExit("Error reading file: %v\n", err)
	}

	if len(fileContents) > 0 {
		panic("Scanner not implemented")
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
