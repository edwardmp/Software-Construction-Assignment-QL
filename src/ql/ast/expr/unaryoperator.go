package expr

import "ql/interfaces"

type UnaryOperator struct {
	Value interfaces.Expr
	interfaces.Expr
}

func (u UnaryOperator) GetValue() interfaces.Expr {
	return u.Value
}
