package binaryoperatorexpr

import "ql/ast/expr"

type And struct {
	BinaryOperator
}

func NewAnd(lhs expr.Expr, rhs expr.Expr) And {
	return And{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (a And) Eval(s interface{}) interface{} {
	return a.Lhs.Eval(s).(bool) && a.Rhs.Eval(s).(bool)
}
