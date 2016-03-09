package binaryoperatorexpr

import (
	"ql/ast/expr"
)

type LT struct {
	Lhs, Rhs expr.Expr
}

func (l LT) Eval(s interface{}) interface{} {
	return l.Lhs.Eval(s).(int) < l.Rhs.Eval(s).(int)
}
