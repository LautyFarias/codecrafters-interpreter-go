package scanning

import (
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/reporter"
	"os"
)

func ScanFile(path string) (scanner *Scanner) {
	file, err := os.Open(path)

	if err != nil {
		reporter.PrintErrorAndExit("Error opening file: %v\n", err)
	}

	scanner = NewScanner(file)
	return
}
