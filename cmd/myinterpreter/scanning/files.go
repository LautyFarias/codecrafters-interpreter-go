package scanning

import (
	"bufio"
	"errors"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/reporter"
	"os"
)

func ScanFile(path string) (scanner *bufio.Scanner, file *os.File) {
	file, err := os.Open(path)

	if err != nil {
		reporter.PrintErrorAndExit("Error opening file: %v\n", err)
	}

	scanner = bufio.NewScanner(file)
	return
}

func GetNextRune(line string, currentIndex int) (rune, error) {
	nextIndex := currentIndex + 1

	if nextIndex >= len(line) {
		return ' ', errors.New("no next rune found")
	}

	return rune(line[nextIndex]), nil
}
