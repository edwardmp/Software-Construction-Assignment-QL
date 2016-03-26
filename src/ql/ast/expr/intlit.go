package expr

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
