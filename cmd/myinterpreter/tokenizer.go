package main

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

	"*": "STAR",
	"/": "SLASH",
	"+": "PLUS",
	"-": "MINUS",
	".": "DOT",
	",": "COMMA",
	";": "SEMI",
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

func tokenize(char string) Token {
	return Token{tokenType: typeByChar[char], lexeme: getLexeme(char), literal: getLiteral(char)}
}
