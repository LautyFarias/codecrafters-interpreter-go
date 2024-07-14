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
	bytes := make([]byte, 1)

	for {
		n, err := reader.Read(bytes)

		if err == io.EOF {
			fmt.Println(tokenize(EOF))
			break
		}

		token, err := tokenize(string(bytes[:n]))

		if err != nil {
			reportCharError(err, n)
		}

		fmt.Println(token)
	}
}
