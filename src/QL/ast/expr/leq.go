package expr

import "ql/ast/visit"

type LEq struct {
	Lhs, Rhs Expr
}

func (l LEq) GetLhs() Expr {
	return l.Lhs
}

func (l LEq) GetRhs() Expr {
	return l.Rhs
}

func (l LEq) Eval() interface{} {
	return l.GetLhs().Eval().(int) <= l.GetRhs().Eval().(int)
}

func (l LEq) Accept(v visit.Visitor) interface{} {
	return v.Visit(l)
}
