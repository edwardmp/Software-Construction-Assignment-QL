package binaryoperatorexpr

import (
	"ql/ast/expr"
	"ql/ast/visit"
)

type Sub struct {
	Lhs, Rhs expr.Expr
}

func (s Sub) GetLhs() expr.Expr {
	return s.Lhs
}

func (s Sub) GetRhs() expr.Expr {
	return s.Rhs
}

func (s Sub) Eval() interface{} {
	return s.GetLhs().Eval().(int) - s.GetRhs().Eval().(int)
}

func (s Sub) Accept(v visit.Visitor, sym interface{}) interface{} {
	return v.Visit(s, sym)
}
