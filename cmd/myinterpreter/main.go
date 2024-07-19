package main

import (
	"fmt"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/reporter"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanning"
	"os"
)

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
	scanner, file := scanning.ScanFile(filename)
	defer file.Close()

	code := NON_ERR_EXIT_CODE
	lineNumber := 0

	for scanner.Scan() {
		lineNumber += 1
		line := scanner.Text()

		skipNext := false

		for index, runeValue := range line {
			if runeValue == scanning.BlankToken || runeValue == scanning.TabToken {
				continue
			}

			if skipNext {
				skipNext = false
				continue
			}

			next := func() string {
				nextIndex := index + 1

				if nextIndex >= len(line) {
					return " "
				}

				return string(line[index+1])
			}()

			char := string(runeValue)
			charset := char + next

			if charset == scanning.CommentToken {
				break
			}

			if scanning.IsToken(charset) {
				char, skipNext = charset, true
			}

			token, err := scanning.Tokenize(char)

			if err != nil {
				reporter.PrintCharError(err, lineNumber)
				code = LEXICAL_ERR_EXIT_CODE

				continue
			}

			fmt.Println(token)
		}
	}

	token, _ := scanning.Tokenize(scanning.EOF)
	fmt.Println(token)

	os.Exit(code)
}
