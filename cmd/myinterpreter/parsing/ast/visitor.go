package ast

import "fmt"

type Visitor interface {
	visitLiteral(*Literal)
}

type PrinterVisitor struct{}

func (v *PrinterVisitor) visitLiteral(literal *Literal) {
	fmt.Println(literal)
}
