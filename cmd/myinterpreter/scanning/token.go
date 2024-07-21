package scanning

import (
	"errors"
	"strings"
)

type Token struct {
	tokenType string
	lexeme    string
	literal   string
}

func (t Token) String() string {
	return t.tokenType + " " + t.lexeme + " " + t.literal
}

func Tokenize(lexeme string) (Token, error) {
	tokenType, err := getType(lexeme)

	if err != nil {
		return Token{}, err
	}

	return Token{tokenType: tokenType, lexeme: getLexeme(lexeme), literal: getLiteral(lexeme)}, nil
}

const EOF = "EOF"

const (
	TabToken   = '	'
	BlankToken = ' '
)

var typeByChar = map[string]string{
	EOF: EOF,
	"(": "LEFT_PAREN",
	")": "RIGHT_PAREN",
	"{": "LEFT_BRACE",
	"}": "RIGHT_BRACE",
	".": "DOT",
	",": "COMMA",
	";": "SEMICOLON",

	// OPERATORS
	// Comparison
	"<":  "LESS",
	"<=": "LESS_EQUAL",
	">":  "GREATER",
	">=": "GREATER_EQUAL",
	"!":  "BANG",
	"!=": "BANG_EQUAL",
	"==": "EQUAL_EQUAL",

	// Math
	"=": "EQUAL",
	"*": "STAR",
	"/": "SLASH",
	"+": "PLUS",
	"-": "MINUS",
}

func getType(char string) (string, error) {
	tokenType, ok := typeByChar[char]

	if !ok {
		if len(char) > 1 && char[0] == '"' && char[len(char)-1] == '"' {
			return "STRING", nil
		}

		if isNumeric(char) {
			return "NUMBER", nil
		}

		return "", errors.New("Unexpected character: " + char)
	}

	return tokenType, nil
}

func getLexeme(char string) string {
	switch char {
	case EOF:
		return ""
	default:
		return char
	}
}

func getLiteral(char string) string {
	if char[0] == '"' && char[len(char)-1] == '"' {
		return strings.ReplaceAll(char, "\"", "")
	}

	if isNumeric(char) {
		if !strings.ContainsRune(char, '.') {
			char += ".0"
		}
		return char
	}

	return "null"
}
