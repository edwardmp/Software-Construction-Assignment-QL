package expr

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
