package expr

import "ql/ast/visit"

type GEq struct {
	Lhs, Rhs Expr
}

func (g GEq) getLhs() Expr {
	return g.Lhs
}

func (g GEq) getRhs() Expr {
	return g.Rhs
}

func (g GEq) Eval() interface{} {
	return g.getLhs().Eval().(int) >= g.getRhs().Eval().(int)
}

func (g GEq) Accept(v visit.Visitor) interface{} {
	return v.Visit(g)
}
