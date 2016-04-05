package expr

import "ql/interfaces"

type BoolLit struct {
	boolValue BoolValue
	Expr
}

func NewBoolLit(value bool) BoolLit {
	return BoolLit{boolValue: NewBoolValue(value), Expr: NewExpr()}
}

func (this BoolLit) Value() interfaces.Value {
	return this.boolValue
}
