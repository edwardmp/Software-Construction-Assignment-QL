package binaryoperatorexpr

import ("ql/ast/visit"
"ql/ast/expr")

type Div struct {
	Lhs, Rhs expr.Expr
}

func (d Div) GetLhs() expr.Expr {
	return d.Lhs
}

func (d Div) GetRhs() expr.Expr {
	return d.Rhs
}

func (d Div) Eval() interface{} {
	return d.GetLhs().Eval().(int) / d.GetRhs().Eval().(int)
}

func (d Div) Accept(v visit.Visitor) interface{} {
	return v.Visit(d)
}
