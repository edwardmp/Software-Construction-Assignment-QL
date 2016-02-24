package binaryoperatorexpr

import (
	"ql/ast/expr"
	"ql/ast/visit"
)

type LT struct {
	Lhs, Rhs expr.Expr
}

func (l LT) GetLhs() expr.Expr {
	return l.Lhs
}

func (l LT) GetRhs() expr.Expr {
	return l.Rhs
}

func (l LT) Eval(s interface{}) interface{} {
	return l.GetLhs().Eval(s).(int) < l.GetRhs().Eval(s).(int)
}

func (l LT) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(l, s)
}
