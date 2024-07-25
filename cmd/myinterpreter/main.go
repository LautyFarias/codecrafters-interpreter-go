package main

import (
	"fmt"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanning"
	"os"
)

const (
	errExitCode        = 1
	lexicalErrExitCode = 65
)

func printErrorAndExit(err string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, err, args...)
	os.Exit(errExitCode)
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		printErrorAndExit("Error opening file: %v\n", err)
	}

	return file
}

func main() {
	if len(os.Args) < 3 {
		printErrorAndExit("Usage: ./your_program.sh <command> <filename>")
	}

	command := os.Args[1]
	filename := os.Args[2]

	switch command {
	case "tokenize":
		file := openFile(filename)
		defer file.Close()

		scanner := scanning.NewScanner(file)
		scanner.Scan()

		if scanner.Error {
			os.Exit(lexicalErrExitCode)
		}
	case "parse":
		printErrorAndExit("Not implemented parse yet")
	default:
		printErrorAndExit("Unknown command: %s\n", command)
	}

}
