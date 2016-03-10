package binaryoperatorexpr

import "ql/ast/expr"

type GT struct {
	BinaryOperator
}

func NewGT(lhs expr.Expr, rhs expr.Expr) GT {
	return GT{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (g GT) Eval(s interface{}) interface{} {
	return g.Lhs.Eval(s).(int) > g.Rhs.Eval(s).(int)
}
