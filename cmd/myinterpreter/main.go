package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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
	bytes := make([]byte, 2)

	code := NON_ERR_EXIT_CODE

	for {
		n, err := reader.Read(bytes)

		if err == io.EOF {
			token, _ := tokenize(EOF)
			fmt.Println(token)

			break
		}

		charset := string(bytes[:n])

		if !isToken(charset) {
			char := charset[:1]

			token, err := tokenize(char)

			if err != nil {
				reportCharError(err, n)
				code = LEXICAL_ERR_EXIT_CODE

				continue
			}

			fmt.Println(token)

			charset = charset[1:]
		}

		token, err := tokenize(charset)

		if err != nil {
			reportCharError(err, n)
			code = LEXICAL_ERR_EXIT_CODE

			continue
		}

		fmt.Println(token)
	}

	quit(code)
}
