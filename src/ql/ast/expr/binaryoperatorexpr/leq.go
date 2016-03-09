package binaryoperatorexpr

import (
	"ql/ast/expr"
)

type LEq struct {
	Lhs, Rhs expr.Expr
}

func (l LEq) Eval(s interface{}) interface{} {
	return l.Lhs.Eval(s).(int) <= l.Rhs.Eval(s).(int)
}