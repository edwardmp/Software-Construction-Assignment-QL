package binaryoperatorexpr

import (
	"ql/ast/expr"
)

type Sub struct {
	Lhs, Rhs expr.Expr
}

func (s Sub) Eval(sy interface{}) interface{} {
	return s.Lhs.Eval(sy).(int) - s.Rhs.Eval(sy).(int)
}