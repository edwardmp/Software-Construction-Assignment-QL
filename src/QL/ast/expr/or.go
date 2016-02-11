package expr

import "ql/ast/visit"

type Or struct {
	Lhs, Rhs Expr
}

func (o Or) getLhs() Expr {
	return o.Lhs
}

func (o Or) getRhs() Expr {
	return o.Rhs
}

func (o Or) Eval() interface{} {
	return o.getLhs().Eval().(bool) || o.getRhs().Eval().(bool)
}

func (o Or) Accept(v visit.Visitor) interface{} {
	return v.Visit(o)
}
