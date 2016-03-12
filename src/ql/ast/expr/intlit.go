package expr

import "fmt"

type IntLit struct {
	Value int
	Expr
}

func NewIntLit(value int, sourceInfo interface{}) IntLit {
	return IntLit{Value: value, Expr: NewExpr(sourceInfo)}
}

func NewIntLitNoSourceInfo(value int) IntLit {
	return NewIntLit(value, nil)
}

func (i IntLit) GetValue() int {
	return i.Value
}

func (i IntLit) String() string {
	return fmt.Sprintf("%d", i.Value)
}
