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

func (this IntLit) GetValue() int {
	return this.Value
}

// FIXME needed for test comparison due to strange behavior. Find better solution.
func (this IntLit) String() string {
	return fmt.Sprintf("%d", this.Value)
}
