package scanning

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Scanner struct {
	source     *os.File
	scanner    *bufio.Scanner
	line       string
	lineNumber int
	chI        int
	Error      bool
}

func NewScanner(source *os.File) *Scanner {
	return &Scanner{source: source, scanner: bufio.NewScanner(source)}
}

func (s *Scanner) next() (rune, error) {
	if s.chI >= len(s.line) {
		return ' ', errors.New("EOL")
	}

	n := s.line[s.chI]
	s.chI++

	return rune(n), nil
}

func (s *Scanner) scanNumber(initial rune) string {
	b := strings.Builder{}
	b.WriteRune(initial)

	defer b.Reset()

	for {
		ch, err := s.next()

		if err != nil {
			return b.String()
		}

		if !unicode.IsNumber(ch) {
			// ch is a letter or string builder already contains '.'
			if ch != '.' || strings.ContainsRune(b.String(), ch) {
				s.chI--
				return b.String()
			}

			// at the moment, ch == '.'

			ch, _ = s.next()
			if !unicode.IsNumber(ch) {
				s.chI--
				return b.String()
			}

			b.WriteRune('.')
		}

		b.WriteRune(ch)
	}
}

func (s *Scanner) scanString(initial rune) (string, error) {
	b := strings.Builder{}
	b.WriteRune(initial)

	defer b.Reset()

	for {
		ch, err := s.next()

		if err != nil {
			return "", errors.New("Unterminated string.")
		}

		b.WriteRune(ch)

		if ch == initial {
			return b.String(), nil
		}
	}
}

func (s *Scanner) scanWord(initial rune) string {
	b := strings.Builder{}
	b.WriteRune(initial)

	defer b.Reset()

	for {
		ch, err := s.next()

		if err != nil {
			return b.String()
		}

		if ch == whitespace {
			s.chI--
			return b.String()
		}

		b.WriteRune(ch)
	}
}

const (
	tab        = '	'
	whitespace = ' '
)

func (s *Scanner) scanLine() {
	var lexeme string

lineIteration:
	for {
		ch, err := s.next()

		if err != nil {
			break
		}

		switch ch {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			lexeme = s.scanNumber(ch)
		case '"':
			lexeme, err = s.scanString(ch)

			if err != nil {
				s.reportError(err)
				continue
			}
		case '<', '>', '=', '!':
			next, _ := s.next()

			lexeme = string(ch)

			if next != '=' {
				s.chI--
				break
			}

			lexeme += string(next)
		case '/':
			next, err := s.next()

			if err != nil {
				lexeme = string(ch)
				break
			}

			if next == '/' {
				break lineIteration
			}

			lexeme = string(ch)
			s.chI--
		case whitespace, tab:
			continue
		default:
			if unicode.IsLetter(ch) || ch == '_' {
				lexeme = s.scanWord(ch)
				break
			}

			lexeme = string(ch)
		}

		s.reportToken(lexeme)
	}
}

func (s *Scanner) Scan() {
	for s.scanner.Scan() {
		s.lineNumber++
		s.line = s.scanner.Text()

		s.scanLine()
		s.chI = 0
	}

	s.reportToken(EOF.String())
}

func (s *Scanner) reportToken(lexeme string) {
	token, err := Tokenize(lexeme)

	if err != nil {
		s.reportError(err)
		return
	}

	fmt.Println(token)
}

func (s *Scanner) reportError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "[line %v] Error: %v\n", s.lineNumber, err)

	s.Error = true
}

func (s *Scanner) Close() {
	_ = s.source.Close()
}
