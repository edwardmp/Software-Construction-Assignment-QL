package expr

import "ql/interfaces"

type IntLit struct {
	intValue IntValue
	Expr
}

func NewIntLit(value int) IntLit {
	return IntLit{intValue: NewIntValue(value), Expr: NewExpr()}
}

func (this IntLit) Value() interfaces.Value {
	return this.intValue
}
