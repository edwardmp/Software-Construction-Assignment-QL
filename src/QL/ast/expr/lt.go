package expr

import "ql/ast/visit"

type LT struct {
	Lhs, Rhs Expr
}

func (l LT) GetLhs() Expr {
	return l.Lhs
}

func (l LT) GetRhs() Expr {
	return l.Rhs
}

func (l LT) Eval() interface{} {
	return l.GetLhs().Eval().(int) < l.GetRhs().Eval().(int)
}

func (l LT) Accept(v visit.Visitor) interface{} {
	return v.Visit(l)
}
