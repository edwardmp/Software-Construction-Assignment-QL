package expr

import "ql/ast/visit"

type Sub struct {
	Lhs, Rhs Expr
}

func (s Sub) getLhs() Expr {
	return s.Lhs
}

func (s Sub) getRhs() Expr {
	return s.Rhs
}

func (s Sub) Eval() interface{} {
	return s.getLhs().Eval().(int) - s.getRhs().Eval().(int)
}

func (s Sub) Accept(v visit.Visitor) interface{} {
	return v.Visit(s)
}
