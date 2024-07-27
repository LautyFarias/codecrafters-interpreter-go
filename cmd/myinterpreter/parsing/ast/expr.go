package ast

import (
	"fmt"
	"math"
	"strconv"
)

type Expr interface {
	accept(Visitor)
}

type Literal struct {
	Value interface{}
}

func (l *Literal) String() string {
	str := fmt.Sprintf("%v", l.Value)

	num, err := strconv.ParseFloat(str, 64)
	if err == nil {
		if num == math.Trunc(num) {
			return fmt.Sprintf("%.1f", num)
		}
		return fmt.Sprintf("%.2f", num)
	}

	return str
}

func (l *Literal) Accept(v Visitor) {
	v.visitLiteral(l)
}
