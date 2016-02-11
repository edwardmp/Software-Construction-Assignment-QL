package expr

import "ql/ast/visit"

type GEq struct {
	Lhs, Rhs Expr
}

func (g GEq) GetLhs() Expr {
	return g.Lhs
}

func (g GEq) GetRhs() Expr {
	return g.Rhs
}

func (g GEq) Eval() interface{} {
	return g.GetLhs().Eval().(int) >= g.GetRhs().Eval().(int)
}

func (g GEq) Accept(v visit.Visitor) interface{} {
	return v.Visit(g)
}
