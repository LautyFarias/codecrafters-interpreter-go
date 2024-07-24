package scanning

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   string
}

func (t Token) String() string {
	return fmt.Sprintf("%s %v %v", t.tokenType, t.lexeme, t.literal)
}

func Tokenize(lexeme string) (Token, error) {
	tt, err := getType(lexeme)

	if err != nil {
		return Token{}, err
	}

	return Token{tokenType: tt, lexeme: getLexeme(lexeme), literal: getLiteral(lexeme, tt)}, nil
}

type TokenType int

var typeNameByTokenType = [...]string{
	EOF:          "EOF",
	LeftParen:    "LEFT_PAREN",
	RightParen:   "RIGHT_PAREN",
	LeftBrace:    "LEFT_BRACE",
	RightBrace:   "RIGHT_BRACE",
	DOT:          "DOT",
	COMMA:        "COMMA",
	SEMICOLON:    "SEMICOLON",
	LESS:         "LESS",
	LessEqual:    "LESS_EQUAL",
	GREATER:      "GREATER",
	GreaterEqual: "GREATER_EQUAL",
	BANG:         "BANG",
	BangEqual:    "BANG_EQUAL",
	EQUAL:        "EQUAL",
	EqualEqual:   "EQUAL_EQUAL",
	STAR:         "STAR",
	SLASH:        "SLASH",
	PLUS:         "PLUS",
	MINUS:        "MINUS",
	STRING:       "STRING",
	NUMBER:       "NUMBER",
}

func (tt TokenType) String() string {
	return typeNameByTokenType[tt]
}

const (
	EOF TokenType = iota

	// CODE BLOCKS

	LeftParen
	RightParen
	LeftBrace
	RightBrace

	// OTHER

	DOT
	COMMA
	SEMICOLON

	// COMPARISON

	LESS
	LessEqual
	GREATER
	GreaterEqual
	BANG
	BangEqual
	EqualEqual

	// MATH

	EQUAL
	STAR
	SLASH
	PLUS
	MINUS

	// COMPLEX

	STRING
	NUMBER
)

func getType(lexeme string) (tt TokenType, err error) {
	switch lexeme {
	case "EOF":
		tt = EOF
	case "(":
		tt = LeftParen
	case ")":
		tt = RightParen
	case "{":
		tt = LeftBrace
	case "}":
		tt = RightBrace
	case ".":
		tt = DOT
	case ",":
		tt = COMMA
	case ";":
		tt = SEMICOLON
	case "<":
		tt = LESS
	case "<=":
		tt = LessEqual
	case ">":
		tt = GREATER
	case ">=":
		tt = GreaterEqual
	case "!":
		tt = BANG
	case "!=":
		tt = BangEqual
	case "==":
		tt = EqualEqual
	case "=":
		tt = EQUAL
	case "*":
		tt = STAR
	case "/":
		tt = SLASH
	case "+":
		tt = PLUS
	case "-":
		tt = MINUS
	default:
		if len(lexeme) > 1 && lexeme[0] == '"' && lexeme[len(lexeme)-1] == '"' {
			tt = STRING
			break
		}

		if _, err := strconv.ParseFloat(lexeme, 64); err == nil {
			tt = NUMBER
			break
		}

		err = errors.New(fmt.Sprintf("Unexpected character: %s", lexeme))
	}

	return tt, err
}

func getLexeme(lexeme string) string {
	if lexeme == EOF.String() {
		return ""
	}

	return lexeme
}

func getLiteral(lexeme string, tt TokenType) string {
	switch tt {
	case STRING:
		return strings.ReplaceAll(lexeme, "\"", "")
	case NUMBER:
		if !strings.ContainsRune(lexeme, '.') {
			lexeme = fmt.Sprintf("%s.0", lexeme)
		}

		return lexeme
	default:
		return "null"
	}
}
