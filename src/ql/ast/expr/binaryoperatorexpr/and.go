package binaryoperatorexpr

import (
	"ql/ast/expr"
)

type And struct {
	Lhs, Rhs expr.Expr
}

func (a And) Eval(s interface{}) interface{} {
	return a.Lhs.(expr.Expr).Eval(s).(bool) && a.Rhs.(expr.Expr).Eval(s).(bool)
}
