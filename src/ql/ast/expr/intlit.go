package expr

import "fmt"

type IntLit struct {
	Value int
	Expr
}

func NewIntLit(value int) IntLit {
	return IntLit{Value: value, Expr: NewExpr()}
}

func (this IntLit) GetValue() int {
	return this.Value
}

// FIXME needed for test comparison due to strange behavior. Find better solution.
func (this IntLit) String() string {
	return fmt.Sprintf("%d", this.Value)
}
