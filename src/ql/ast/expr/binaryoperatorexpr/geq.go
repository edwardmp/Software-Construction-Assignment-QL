package binaryoperatorexpr

import (
	"ql/ast/expr"
)

type GEq struct {
	Lhs, Rhs expr.Expr
}

func (g GEq) Eval(s interface{}) interface{} {
	return g.Lhs.Eval(s).(int) >= g.Rhs.Eval(s).(int)
}
