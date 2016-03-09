package binaryoperatorexpr

import (
	"ql/ast/expr"
)

type Add struct {
	Lhs, Rhs expr.Expr
}

func (a Add) Eval(s interface{}) interface{} {
	return a.Lhs.Eval(s).(int) + a.Rhs.Eval(s).(int)
}
