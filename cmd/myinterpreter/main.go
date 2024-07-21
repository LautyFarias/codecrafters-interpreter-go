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

func checkArguments(args []string) {
	if len(args) < 3 {
		printErrorAndExit("Usage: ./your_program.sh tokenize <filename>")
	}

	command := args[1]

	if command != "tokenize" {
		printErrorAndExit("Unknown command: %s\n", command)
	}
}

func main() {
	checkArguments(os.Args)

	filename := os.Args[2]
	file, err := os.Open(filename)

	if err != nil {
		printErrorAndExit("Error opening file: %v\n", err)
	}

	scanner := scanning.NewScanner(file)
	defer scanner.Close()

	scanner.Scan()

	if scanner.Error {
		os.Exit(lexicalErrExitCode)
	}
}
