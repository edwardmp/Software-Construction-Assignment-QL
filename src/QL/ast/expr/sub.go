package expr

import "ql/ast/visit"

type Sub struct {
	Lhs, Rhs Expr
}

func (s Sub) GetLhs() Expr {
	return s.Lhs
}

func (s Sub) GetRhs() Expr {
	return s.Rhs
}

func (s Sub) Eval() interface{} {
	return s.GetLhs().Eval().(int) - s.GetRhs().Eval().(int)
}

func (s Sub) Accept(v visit.Visitor) interface{} {
	return v.Visit(s)
}
