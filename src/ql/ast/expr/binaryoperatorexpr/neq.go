package binaryoperatorexpr

import (
	"ql/ast/expr"
	"ql/ast/visit"
)

type NEq struct {
	Lhs, Rhs expr.Expr
}

func (n NEq) GetLhs() expr.Expr {
	return n.Lhs
}

func (n NEq) GetRhs() expr.Expr {
	return n.Rhs
}

func (n NEq) Eval(s interface{}) interface{} {
	switch n.Lhs.Eval(s).(type) {
	case int:
		return n.GetLhs().Eval(s).(int) != n.GetRhs().Eval(s).(int)
	case bool:
		return n.GetLhs().Eval(s).(bool) != n.GetRhs().Eval(s).(bool)
	default:
		panic("NEq error: comparing unknown types")
	}
}

func (n NEq) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(n, s)
}
