package expr

import "ql/interfaces"

type Div struct {
	BinaryOperator
}

func NewDiv(lhs interfaces.Expr, rhs interfaces.Expr) Div {
	return Div{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (d Div) Eval(s interface{}) interface{} {
	return d.Lhs.Eval(s).(int) / d.Rhs.Eval(s).(int)
}
