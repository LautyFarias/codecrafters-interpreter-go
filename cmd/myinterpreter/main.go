package main

import (
	"fmt"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/reporter"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanning"
	"os"
)

const lexicalErrExitCode = 65

func checkArguments(args []string) {
	if len(args) < 3 {
		reporter.PrintErrorAndExit("Usage: ./your_program.sh tokenize <filename>")
	}

	command := args[1]

	if command != "tokenize" {
		reporter.PrintErrorAndExit("Unknown command: %s\n", command)
	}
}

func main() {
	checkArguments(os.Args)

	filename := os.Args[2]
	scanner := scanning.ScanFile(filename)
	defer scanner.Close()

	scanner.Scan()

	token, _ := scanning.Tokenize(scanning.EOF)
	fmt.Println(token)

	if scanner.Error {
		os.Exit(lexicalErrExitCode)
	}
}
