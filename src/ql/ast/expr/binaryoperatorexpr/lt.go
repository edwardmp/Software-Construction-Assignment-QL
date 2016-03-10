package binaryoperatorexpr

import "ql/ast/expr"

type LT struct {
	BinaryOperator
}

func NewLT(lhs expr.Expr, rhs expr.Expr) LT {
	return LT{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (l LT) Eval(s interface{}) interface{} {
	return l.Lhs.Eval(s).(int) < l.Rhs.Eval(s).(int)
}
