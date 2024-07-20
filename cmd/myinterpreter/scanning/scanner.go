package scanning

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/reporter"
	"os"
	"strings"
)

type Scanner struct {
	source        *os.File
	scanner       *bufio.Scanner
	stringBuilder *strings.Builder
	line          string
	lineNumber    int
	charIndex     int
	Error         bool
}

func NewScanner(source *os.File) *Scanner {
	return &Scanner{
		source:        source,
		scanner:       bufio.NewScanner(source),
		stringBuilder: &strings.Builder{},
	}
}

func (s *Scanner) next() (rune, error) {
	nextIndex := s.charIndex + 1

	if nextIndex >= len(s.line) {
		return ' ', errors.New("no next rune found")
	}

	return rune(s.line[nextIndex]), nil
}

func (s *Scanner) isBuildingString() bool {
	return s.stringBuilder.Len() > 0 && s.stringBuilder.String()[0] == '"'
}

func (s *Scanner) isBuildingNumber() bool {
	return s.stringBuilder.Len() > 0 && s.stringBuilder.String()[0] != '"'
}

func (s *Scanner) Scan() {
	for s.scanner.Scan() {
		s.lineNumber++

		s.line = s.scanner.Text()

		var lexeme string
		var char rune
		skipChar := false

	lineIteration:
		for s.charIndex, char = range s.line {
			if skipChar {
				skipChar = false
				continue
			}

			if s.isBuildingString() && char != '"' {
				s.stringBuilder.WriteRune(char)
				continue
			}

			switch char {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				s.stringBuilder.WriteRune(char)
				continue
			case '.':
				if s.isBuildingNumber() {
					str := s.stringBuilder.String()

					if strings.ContainsRune(str, char) {
						s.reportToken(str)
						s.stringBuilder.Reset()

						lexeme = string(char)
						break
					}

					s.stringBuilder.WriteRune(char)
					continue
				}

				lexeme = string(char)
			case '"':
				s.stringBuilder.WriteRune(char)

				if s.stringBuilder.Len() == 1 {
					continue
				}

				if s.isBuildingString() {
					lexeme = s.stringBuilder.String()
					s.stringBuilder.Reset()
				}
			case '<', '>', '=', '!':
				next, _ := s.next()

				lexeme = string(char)

				if next == '=' {
					lexeme += string(next)
					skipChar = true
				}
			case '/':
				next, err := s.next()

				if err != nil {
					lexeme = string(char)
				}

				if next == '/' {
					break lineIteration
				}
			case BlankToken, TabToken:
				continue
			default:
				lexeme = string(char)
			}

			s.reportToken(lexeme)
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
	reporter.PrintErrorAtLine(err, s.lineNumber)
	s.Error = true
}

func (s *Scanner) Close() {
	s.source.Close()
}
