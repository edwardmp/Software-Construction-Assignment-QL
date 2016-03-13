package expr

import "ql/interfaces"

type UnaryOperator struct {
	Value interfaces.Expr
	Expr
}

func NewUnaryOperator(value interfaces.Expr, sourceInfo interface{}) UnaryOperator {
	return UnaryOperator{value, NewExpr(sourceInfo)}
}

func (this UnaryOperator) GetValue() interfaces.Expr {
	return this.Value
}
