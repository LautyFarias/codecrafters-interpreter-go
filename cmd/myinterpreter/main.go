package main

import (
	"fmt"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/reporter"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanning"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		reporter.PrintErrorAndExit("Usage: ./your_program.sh tokenize <filename>")
	}

	command := os.Args[1]

	if command != "tokenize" {
		reporter.PrintErrorAndExit("Unknown command: %s\n", command)
	}

	filename := os.Args[2]

	scanner, file := scanning.ScanFile(filename)
	defer file.Close()

	lineNumber := 0
	code := NON_ERR_EXIT_CODE

	for scanner.Scan() {
		lineNumber += 1
		line := scanner.Text()

		skipNext := false

		for index, runeValue := range line {
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

			if char == scanning.BlankToken || char == scanning.TabToken {
				continue
			}

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
