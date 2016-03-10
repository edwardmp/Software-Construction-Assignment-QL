package binaryoperatorexpr

import "ql/ast/expr"

type LEq struct {
	BinaryOperator
}

func NewLEq(lhs expr.Expr, rhs expr.Expr) LEq {
	return LEq{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (l LEq) Eval(s interface{}) interface{} {
	return l.Lhs.Eval(s).(int) <= l.Rhs.Eval(s).(int)
}
