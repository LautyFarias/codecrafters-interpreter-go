package main

import (
	"fmt"
	"os"
)

const (
	NON_ERR_EXIT_CODE     = 0
	ERR_EXIT_CODE         = 1
	LEXICAL_ERR_EXIT_CODE = 65
)

func quit(code int) {
	os.Exit(code)
}

func printError(err string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, err, args...)
}

func printErrorAndExit(err string, args ...any) {
	printError(err, args...)
	quit(ERR_EXIT_CODE)
}

func reportCharError(err error, line int) {
	printError("[line %v] Error: %v", line, err)
}
