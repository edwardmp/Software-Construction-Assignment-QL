package binaryoperatorexpr

import (
	"ql/ast/expr"
	"ql/ast/visit"
)

type LEq struct {
	Lhs, Rhs expr.Expr
}

func (l LEq) GetLhs() expr.Expr {
	return l.Lhs
}

func (l LEq) GetRhs() expr.Expr {
	return l.Rhs
}

func (l LEq) Eval(s interface{}) interface{} {
	return l.GetLhs().Eval(s).(int) <= l.GetRhs().Eval(s).(int)
}

func (l LEq) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(l, s)
}
