package binaryoperatorexpr

import (
	"ql/ast/expr"
)

type Mul struct {
	Lhs, Rhs expr.Expr
}

func (m Mul) Eval(s interface{}) interface{} {
	return m.Lhs.Eval(s).(int) * m.Rhs.Eval(s).(int)
}
