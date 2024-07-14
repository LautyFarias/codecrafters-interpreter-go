package main

import (
	"fmt"
	"os"
)

func printError(err string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, err, args...)
}

func printErrorAndExit(err string, args ...any) {
	printError(err, args)
	os.Exit(1)
}

func reportCharError(err error, line int) {
	printError("[line %v] Error: %v", line, err)
}
