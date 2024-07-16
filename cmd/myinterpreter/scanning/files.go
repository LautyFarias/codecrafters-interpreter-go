package scanning

import (
	"bufio"
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
