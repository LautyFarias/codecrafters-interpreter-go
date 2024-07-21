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
	source        *os.File
	scanner       *bufio.Scanner
	stringBuilder *strings.Builder
	line          string
	lineNumber    int
	chI           int
	Error         bool
}

func NewScanner(source *os.File) *Scanner {
	return &Scanner{
		source:        source,
		scanner:       bufio.NewScanner(source),
		stringBuilder: &strings.Builder{},
	}
}

func (s *Scanner) isBuildingString() bool {
	return s.stringBuilder.Len() > 0 && s.stringBuilder.String()[0] == '"'
}

func (s *Scanner) isBuildingNumber() bool {
	return s.stringBuilder.Len() > 0 && s.stringBuilder.String()[0] != '"'
}

func (s *Scanner) next() (rune, error) {
	if s.chI >= len(s.line) {
		return ' ', errors.New("EOL")
	}

	n := s.line[s.chI]
	s.chI++

	return rune(n), nil
}

func (s *Scanner) scanLine() {
	var lexeme string

lineIteration:
	for {
		ch, err := s.next()

		if err != nil {
			break
		}

		if s.isBuildingString() && ch != '"' {
			s.stringBuilder.WriteRune(ch)
			continue
		}

		switch ch {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			s.stringBuilder.WriteRune(ch)
			continue
		case '.':
			if s.isBuildingNumber() {
				str := s.stringBuilder.String()

				if strings.ContainsRune(str, ch) {
					s.reportToken(str)
					s.stringBuilder.Reset()

					lexeme = string(ch)
					break
				}

				next, err := s.next()

				if err == nil && unicode.IsNumber(next) {
					s.stringBuilder.WriteRune(ch)
					continue
				}

				s.reportToken(s.stringBuilder.String())
				s.stringBuilder.Reset()
			}

			lexeme = string(ch)
		case '"':
			s.stringBuilder.WriteRune(ch)

			if s.stringBuilder.Len() == 1 {
				continue
			}

			if s.isBuildingString() {
				lexeme = s.stringBuilder.String()
				s.stringBuilder.Reset()
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
		case BlankToken, TabToken:
			continue
		default:
			lexeme = string(ch)
		}

		if s.isBuildingNumber() {
			s.reportToken(s.stringBuilder.String())
			s.stringBuilder.Reset()
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

		if s.isBuildingNumber() {
			s.reportToken(s.stringBuilder.String())
			s.stringBuilder.Reset()
		}

		if s.isBuildingString() {
			s.reportError(errors.New("Unterminated string."))
		}
	}
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
	_, _ = fmt.Fprintf(os.Stderr, "[line %v] Error: %v\n", s.line, err)

	s.Error = true
}

func (s *Scanner) Close() {
	s.source.Close()
}
