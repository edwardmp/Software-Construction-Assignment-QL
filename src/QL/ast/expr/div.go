package expr

import "ql/ast/visit"

type Div struct {
	Lhs, Rhs Expr
}

func (d Div) GetLhs() Expr {
	return d.Lhs
}

func (d Div) GetRhs() Expr {
	return d.Rhs
}

func (d Div) Eval() interface{} {
	return d.GetLhs().Eval().(int) / d.GetRhs().Eval().(int)
}

func (d Div) Accept(v visit.Visitor) interface{} {
	return v.Visit(d)
}
