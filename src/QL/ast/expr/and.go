package expr

import "ql/ast/visit"

type And struct {
	Lhs, Rhs Expr
}

func (a And) GetLhs() Expr {
	return a.Lhs
}

func (a And) GetRhs() Expr {
	return a.Rhs
}

func (a And) Eval() interface{} {
	return a.GetLhs().(Expr).Eval().(bool) && a.GetRhs().(Expr).Eval().(bool)
}

func (a And) Accept(v visit.Visitor) interface{} {
	return v.Visit(a)
}
