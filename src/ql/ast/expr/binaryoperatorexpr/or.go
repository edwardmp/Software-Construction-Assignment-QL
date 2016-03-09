package binaryoperatorexpr

import (
	"ql/ast/expr"
)

type Or struct {
	Lhs, Rhs expr.Expr
}

func (o Or) Eval(s interface{}) interface{} {
	return o.Lhs.Eval(s).(bool) || o.Rhs.Eval(s).(bool)
}

