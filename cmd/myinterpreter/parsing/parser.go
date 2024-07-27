package parsing

import (
	"bufio"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/parsing/ast"
	"os"
)

type Scanner struct {
	source  *os.File
	scanner *bufio.Scanner
	line    string
}

func NewScanner(source *os.File) *Scanner {
	return &Scanner{source: source, scanner: bufio.NewScanner(source)}
}

func (s *Scanner) Scan() {
	for s.scanner.Scan() {
		s.line = s.scanner.Text()

		expr := ast.Literal{Value: s.line}
		expr.Accept(&ast.PrinterVisitor{})
	}
}
