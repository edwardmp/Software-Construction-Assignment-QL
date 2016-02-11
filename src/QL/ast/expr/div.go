package expr

import "ql/ast/visit"

type Div struct {
	Lhs, Rhs Expr
}

func (d Div) getLhs() Expr {
	return d.Lhs
}

func (d Div) getRhs() Expr {
	return d.Rhs
}

func (d Div) Eval() interface{} {
	return d.getLhs().Eval().(int) / d.getRhs().Eval().(int)
}

func (d Div) Accept(v visit.Visitor) interface{} {
	return v.Visit(d)
}
