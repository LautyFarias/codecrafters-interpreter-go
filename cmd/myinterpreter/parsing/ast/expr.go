package ast

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Expr interface {
	accept(Visitor)
}

type Literal struct {
	Value interface{}
}

func (l *Literal) String() string {
	literal := fmt.Sprintf("%v", l.Value)

	num, err := strconv.ParseFloat(literal, 64)
	if err == nil {
		if num == math.Trunc(num) {
			return fmt.Sprintf("%.1f", num)
		}
		return fmt.Sprintf("%.2f", num)
	}

	literal = strings.ReplaceAll(literal, "\"", "")

	return literal
}

func (l *Literal) Accept(v Visitor) {
	v.visitLiteral(l)
}
