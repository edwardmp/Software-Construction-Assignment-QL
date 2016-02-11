package expr

import "ql/ast/visit"

type Add struct {
	Lhs, Rhs Expr
}

func (a Add) GetLhs() Expr {
	return a.Lhs
}

func (a Add) GetRhs() Expr {
	return a.Rhs
}

func (a Add) Eval() interface{} {
	return a.GetLhs().Eval().(int) + a.GetRhs().Eval().(int)
}

func (a Add) Accept(v visit.Visitor) interface{} {
	return v.Visit(a)
}
