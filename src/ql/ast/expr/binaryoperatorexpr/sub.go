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

func (s Sub) Eval(sy interface{}) interface{} {
	return s.GetLhs().Eval(sy).(int) - s.GetRhs().Eval(sy).(int)
}

func (s Sub) Accept(v visit.Visitor, sym interface{}) interface{} {
	return v.Visit(s, sym)
}
