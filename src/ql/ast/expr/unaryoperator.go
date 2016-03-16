package expr

import "ql/interfaces"

type UnaryOperator struct {
	Value interfaces.Expr
	Expr
}

func NewUnaryOperator(value interfaces.Expr, sourceInfo interface{}) UnaryOperator {
	return UnaryOperator{value, NewExpr(sourceInfo)}
}

func (u UnaryOperator) GetValue() interfaces.Expr {
	return u.Value
}
