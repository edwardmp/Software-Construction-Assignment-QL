package expr

import "ql/ast/visit"

type GT struct {
	Lhs, Rhs Expr
}

func (g GT) getLhs() Expr {
	return g.Lhs
}

func (g GT) getRhs() Expr {
	return g.Rhs
}

func (g GT) Eval() interface{} {
	return g.getLhs().Eval().(int) > g.getRhs().Eval().(int)
}

func (g GT) Accept(v visit.Visitor) interface{} {
	return v.Visit(g)
}
