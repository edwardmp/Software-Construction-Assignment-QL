package expr

import "ql/ast/visit"

type LT struct {
	Lhs, Rhs Expr
}

func (l LT) getLhs() Expr {
	return l.Lhs
}

func (l LT) getRhs() Expr {
	return l.Rhs
}

func (l LT) Eval() interface{} {
	return l.getLhs().Eval().(int) < l.getRhs().Eval().(int)
}

func (l LT) Accept(v visit.Visitor) interface{} {
	return v.Visit(l)
}
