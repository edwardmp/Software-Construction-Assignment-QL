package expr

import "ql/ast/visit"

type Or struct {
	Lhs, Rhs Expr
}

func (o Or) GetLhs() Expr {
	return o.Lhs
}

func (o Or) GetRhs() Expr {
	return o.Rhs
}

func (o Or) Eval() interface{} {
	return o.GetLhs().Eval().(bool) || o.GetRhs().Eval().(bool)
}

func (o Or) Accept(v visit.Visitor) interface{} {
	return v.Visit(o)
}
