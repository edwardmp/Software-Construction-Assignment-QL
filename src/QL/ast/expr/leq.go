package expr

import "ql/ast/visit"

type LEq struct {
	Lhs, Rhs Expr
}

func (l LEq) getLhs() Expr {
	return l.Lhs
}

func (l LEq) getRhs() Expr {
	return l.Rhs
}

func (l LEq) Eval() interface{} {
	return l.getLhs().Eval().(int) <= l.getRhs().Eval().(int)
}

func (l LEq) Accept(v visit.Visitor) interface{} {
	return v.Visit(l)
}
