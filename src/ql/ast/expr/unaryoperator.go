package expr

import "ql/interfaces"

type UnaryOperator struct {
	Value interfaces.Expr
	Expr
}

func NewUnaryOperator(value interfaces.Expr) UnaryOperator {
	return UnaryOperator{value, NewExpr()}
}

func (this UnaryOperator) GetValue() interfaces.Expr {
	return this.Value
}
