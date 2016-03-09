package binaryoperatorexpr

import (
	"ql/ast/expr"
)

type Div struct {
	Lhs, Rhs expr.Expr
}

func (d Div) Eval(s interface{}) interface{} {
	return d.Lhs.Eval(s).(int) / d.Rhs.Eval(s).(int)
}
