package binaryoperatorexpr

import (
	"ql/ast/expr"
	"ql/ast/visit"
)

type Div struct {
	Lhs, Rhs expr.Expr
}

func (d Div) GetLhs() expr.Expr {
	return d.Lhs
}

func (d Div) GetRhs() expr.Expr {
	return d.Rhs
}

func (d Div) Eval(s interface{}) interface{} {
	return d.GetLhs().Eval(s).(int) / d.GetRhs().Eval(s).(int)
}

func (d Div) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(d, s)
}
