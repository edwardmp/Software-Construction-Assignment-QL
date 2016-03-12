package expr

import "fmt"

type BoolLit struct {
	Value bool
	Expr
}

func NewBoolLit(value bool, sourceInfo interface{}) BoolLit {
	return BoolLit{Value: value, Expr: NewExpr(sourceInfo)}
}

func NewBoolLitNoSourceInfo(value bool) BoolLit {
	return NewBoolLit(value, nil)
}

func (b BoolLit) GetValue() bool {
	return b.Value
}

func (b BoolLit) String() string {
	return fmt.Sprintf("%t", b.Value)
}
