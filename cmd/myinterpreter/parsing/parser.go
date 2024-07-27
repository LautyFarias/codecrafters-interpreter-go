package parsing

import (
	"bufio"
	"fmt"
	"os"
)

type Scanner struct {
	source  *os.File
	scanner *bufio.Scanner
	lno     int
	l       string
}

func NewScanner(source *os.File) *Scanner {
	return &Scanner{source: source, scanner: bufio.NewScanner(source)}
}

func (s *Scanner) Scan() {
	for s.scanner.Scan() {
		s.lno++
		s.l = s.scanner.Text()

		fmt.Println(s.l)
	}
}
