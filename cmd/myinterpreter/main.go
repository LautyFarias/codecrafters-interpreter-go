package main

import (
	"fmt"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/reporter"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanning"
	"os"
	"strings"
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

	stringBuilder := strings.Builder{}
	buildingString := false

	for scanner.Scan() {
		lineNumber += 1
		line := scanner.Text()

		var lexeme string
		skipChar := false

	lineIteration:
		for index, char := range line {
			if skipChar {
				skipChar = false
				continue
			}

			switch char {
			case '"':
				stringBuilder.WriteRune(char)

				if stringBuilder.Len() == 1 {
					buildingString = true
					continue
				}

				if buildingString {
					buildingString = false
					lexeme = stringBuilder.String()
					stringBuilder.Reset()
				}
			case '<', '>', '=', '!':
				next, _ := scanning.GetNextRune(line, index)

				lexeme = string(char)

				if next == '=' {
					lexeme += string(next)
					skipChar = true
				}

			case '/':
				next, err := scanning.GetNextRune(line, index)

				if err != nil {
					lexeme = string(char)
				}

				if next == '/' {
					break lineIteration
				}
			case scanning.BlankToken, scanning.TabToken:
				if buildingString {
					stringBuilder.WriteRune(char)
				}

				continue
			default:
				if buildingString {
					stringBuilder.WriteRune(char)
					continue
				}

				lexeme = string(char)
			}

			token, err := scanning.Tokenize(lexeme)

			if err != nil {
				reporter.PrintCharError(err, lineNumber)
				code = LEXICAL_ERR_EXIT_CODE

				continue
			}

			fmt.Println(token)
		}

		if buildingString {
			fmt.Printf("[line %v] Error: Unterminated string.\n", lineNumber)
		}
	}

	token, _ := scanning.Tokenize(scanning.EOF)
	fmt.Println(token)

	os.Exit(code)
}
