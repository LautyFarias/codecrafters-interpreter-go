package reporter

import (
	"fmt"
	"os"
)

func printError(err string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, err, args...)
}

func PrintErrorAndExit(err string, args ...any) {
	printError(err, args...)
	os.Exit(1)
}

func PrintErrorAtLine(err error, line int) {
	printError("[line %v] Error: %v\n", line, err)
}
