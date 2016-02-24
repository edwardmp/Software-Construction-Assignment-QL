package binaryoperatorexpr

import (
	"ql/ast/expr"
	"ql/ast/visit"
)

type Mul struct {
	Lhs, Rhs expr.Expr
}

func (m Mul) GetLhs() expr.Expr {
	return m.Lhs
}

func (m Mul) GetRhs() expr.Expr {
	return m.Rhs
}

func (m Mul) Eval(s interface{}) interface{} {
	return m.GetLhs().Eval(s).(int) * m.GetRhs().Eval(s).(int)
}

func (m Mul) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(m, s)
}
