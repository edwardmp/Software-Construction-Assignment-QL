package binaryoperatorexpr

import (
	"ql/ast/expr"
	"ql/ast/visit"
)

type Add struct {
	Lhs, Rhs expr.Expr
}

func (a Add) GetLhs() expr.Expr {
	return a.Lhs
}

func (a Add) GetRhs() expr.Expr {
	return a.Rhs
}

func (a Add) Eval(s interface{}) interface{} {
	return a.GetLhs().Eval(s).(int) + a.GetRhs().Eval(s).(int)
}

func (a Add) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(a, s)
}
