package binaryoperatorexpr

import "ql/ast/expr"

type Add struct {
	BinaryOperator
}

func NewAdd(lhs expr.Expr, rhs expr.Expr) Add {
	return Add{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (a Add) Eval(s interface{}) interface{} {
	return a.Lhs.Eval(s).(int) + a.Rhs.Eval(s).(int)
}
