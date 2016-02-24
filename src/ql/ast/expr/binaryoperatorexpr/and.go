package binaryoperatorexpr

import (
	"ql/ast/expr"
	"ql/ast/visit"
)

type And struct {
	Lhs, Rhs expr.Expr
}

func (a And) GetLhs() expr.Expr {
	return a.Lhs
}

func (a And) GetRhs() expr.Expr {
	return a.Rhs
}

func (a And) Eval(s interface{}) interface{} {
	return a.GetLhs().(expr.Expr).Eval(s).(bool) && a.GetRhs().(expr.Expr).Eval(s).(bool)
}

func (a And) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(a, s)
}
