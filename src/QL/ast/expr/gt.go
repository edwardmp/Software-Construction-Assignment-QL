package expr

import "ql/ast/visit"

type GT struct {
	Lhs, Rhs Expr
}

func (g GT) GetLhs() Expr {
	return g.Lhs
}

func (g GT) GetRhs() Expr {
	return g.Rhs
}

func (g GT) Eval() interface{} {
	return g.GetLhs().Eval().(int) > g.GetRhs().Eval().(int)
}

func (g GT) Accept(v visit.Visitor) interface{} {
	return v.Visit(g)
}
