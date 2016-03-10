package expr

import "ql/interfaces"

type Sub struct {
	BinaryOperator
}

func NewSub(lhs interfaces.Expr, rhs interfaces.Expr) Sub {
	return Sub{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (s Sub) Eval(sy interface{}) interface{} {
	return s.Lhs.Eval(sy).(int) - s.Rhs.Eval(sy).(int)
}
