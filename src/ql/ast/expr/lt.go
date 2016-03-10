package expr

import "ql/interfaces"

type LT struct {
	BinaryOperator
}

func NewLT(lhs interfaces.Expr, rhs interfaces.Expr) LT {
	return LT{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (l LT) Eval(s interface{}) interface{} {
	return l.Lhs.Eval(s).(int) < l.Rhs.Eval(s).(int)
}
