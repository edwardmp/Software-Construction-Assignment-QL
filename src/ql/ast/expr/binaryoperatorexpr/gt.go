package binaryoperatorexpr

import ("ql/ast/visit"
"ql/ast/expr")

type GT struct {
	Lhs, Rhs expr.Expr
}

func (g GT) GetLhs() expr.Expr {
	return g.Lhs
}

func (g GT) GetRhs() expr.Expr {
	return g.Rhs
}

func (g GT) Eval() interface{} {
	return g.GetLhs().Eval().(int) > g.GetRhs().Eval().(int)
}

func (g GT) Accept(v visit.Visitor) interface{} {
	return v.Visit(g)
}
