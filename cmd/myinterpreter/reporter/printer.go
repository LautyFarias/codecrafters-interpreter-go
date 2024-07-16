package reporter

import (
	"fmt"
	"os"
)

func PrintError(err string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, err, args...)
}

func PrintErrorAndExit(err string, args ...any) {
	PrintError(err, args...)
	os.Exit(1)
}

func PrintCharError(err error, line int) {
	PrintError("[line %v] Error: %v\n", line, err)
}
