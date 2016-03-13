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

func (this BoolLit) GetValue() bool {
	return this.Value
}

// FIXME needed for test comparison due to strange behavior. Find better solution.
func (this BoolLit) String() string {
	return fmt.Sprintf("%t", this.Value)
}
