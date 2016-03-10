package expr

import "ql/interfaces"

type Add struct {
	BinaryOperator
}

func NewAdd(lhs interfaces.Expr, rhs interfaces.Expr) Add {
	return Add{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (a Add) Eval(s interface{}) interface{} {
	return a.Lhs.Eval(s).(int) + a.Rhs.Eval(s).(int)
}
