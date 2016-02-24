package binaryoperatorexpr

import (
	"ql/ast/expr"
	"ql/ast/visit"
)

type GEq struct {
	Lhs, Rhs expr.Expr
}

func (g GEq) GetLhs() expr.Expr {
	return g.Lhs
}

func (g GEq) GetRhs() expr.Expr {
	return g.Rhs
}

func (g GEq) Eval(s interface{}) interface{} {
	return g.GetLhs().Eval(s).(int) >= g.GetRhs().Eval(s).(int)
}

func (g GEq) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(g, s)
}
