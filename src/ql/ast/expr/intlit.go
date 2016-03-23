package expr

import "fmt"

type IntLit struct {
	value int
	Expr
}

func NewIntLit(value int) IntLit {
	return IntLit{value: value, Expr: NewExpr()}
}

func (this IntLit) Value() int {
	return this.value
}

// FIXME needed for test comparison due to strange behavior. Find better solution.
func (this IntLit) String() string {
	return fmt.Sprintf("%d", this.value)
}
