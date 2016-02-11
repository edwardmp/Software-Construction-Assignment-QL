package expr

import "ql/ast/visit"

type And struct {
	Lhs, Rhs Expr
}

func (a And) getLhs() Expr {
	return a.Lhs
}

func (a And) getRhs() Expr {
	return a.Rhs
}

func (a And) Eval() interface{} {
	return a.getLhs().(Expr).Eval().(bool) && a.getRhs().(Expr).Eval().(bool)
}

func (a And) Accept(v visit.Visitor) interface{} {
	return v.Visit(a)
}
