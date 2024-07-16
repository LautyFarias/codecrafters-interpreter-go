package scanning

import "errors"

const EOF = "EOF"

type Token struct {
	tokenType string
	lexeme    string
	literal   string
}

func (t Token) String() string {
	return t.tokenType + " " + t.lexeme + " " + t.literal
}

var typeByChar = map[string]string{
	EOF: EOF,
	"(": "LEFT_PAREN",
	")": "RIGHT_PAREN",
	"{": "LEFT_BRACE",
	"}": "RIGHT_BRACE",

	"==": "EQUAL_EQUAL",
	"=":  "EQUAL",
	"*":  "STAR",
	"/":  "SLASH",
	"+":  "PLUS",
	"-":  "MINUS",
	".":  "DOT",
	",":  "COMMA",
	";":  "SEMICOLON",
}

func IsToken(char string) (ok bool) {
	_, ok = typeByChar[char]
	return
}

func getType(char string) (string, error) {
	tokenType, ok := typeByChar[char]

	if !ok {
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
	switch char {
	default:
		return "null"
	}
}

func Tokenize(char string) (Token, error) {
	tokenType, err := getType(char)

	if err != nil {
		return Token{}, err
	}

	return Token{tokenType: tokenType, lexeme: getLexeme(char), literal: getLiteral(char)}, nil
}
