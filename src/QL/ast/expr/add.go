package expr

import "ql/ast/visit"

type Add struct {
	Lhs, Rhs Expr
}

func (a Add) getLhs() Expr {
	return a.Lhs
}

func (a Add) getRhs() Expr {
	return a.Rhs
}

func (a Add) Eval() interface{} {
	return a.getLhs().Eval().(int) + a.getRhs().Eval().(int)
}

func (a Add) Accept(v visit.Visitor) interface{} {
	return v.Visit(a)
}
