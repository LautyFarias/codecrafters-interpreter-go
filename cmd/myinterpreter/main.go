package main

import (
	"bufio"
	"fmt"
	"os"
)

func scanFile(path string) (scanner *bufio.Scanner, file *os.File) {
	file, err := os.Open(path)

	if err != nil {
		printErrorAndExit("Error opening file: %v\n", err)
	}

	scanner = bufio.NewScanner(file)
	return
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

	scanner, file := scanFile(filename)
	defer file.Close()

	lineNumber := 0
	code := NON_ERR_EXIT_CODE

	for scanner.Scan() {
		lineNumber += 1
		line := scanner.Text()

		skipNext := false

		for index, runeValue := range line {
			char := string(runeValue)

			if skipNext {
				skipNext = false
				continue
			}

			next := func() string {
				if len(line[index:]) > 2 {
					return string(line[index+1])
				}

				return " "
			}()

			charset := char + next

			var token Token
			var err error

			if isToken(charset) {
				token, err = tokenize(charset)
				skipNext = true
			} else {
				token, err = tokenize(char)
			}

			if err != nil {
				reportCharError(err, lineNumber)
				code = LEXICAL_ERR_EXIT_CODE
			} else {
				fmt.Println(token)
			}
		}
	}

	token, _ := tokenize(EOF)
	fmt.Println(token)

	quit(code)
}
