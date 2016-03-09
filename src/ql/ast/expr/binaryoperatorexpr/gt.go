package binaryoperatorexpr

import (
	"ql/ast/expr"
)

type GT struct {
	Lhs, Rhs expr.Expr
}

func (g GT) Eval(s interface{}) interface{} {
	return g.Lhs.Eval(s).(int) > g.Rhs.Eval(s).(int)
}

