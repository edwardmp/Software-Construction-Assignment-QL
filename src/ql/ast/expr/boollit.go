package expr

import "fmt"

type BoolLit struct {
	value bool
	Expr
}

func NewBoolLit(value bool) BoolLit {
	return BoolLit{value: value, Expr: NewExpr()}
}

func (this BoolLit) Value() bool {
	return this.value
}

// FIXME needed for test comparison due to strange behavior. Find better solution.
func (this BoolLit) String() string {
	return fmt.Sprintf("%T", this.value)
}
