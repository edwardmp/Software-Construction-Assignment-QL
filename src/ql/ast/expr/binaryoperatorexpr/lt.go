package binaryoperatorexpr

import ("ql/ast/visit"
"ql/ast/expr")

type LT struct {
	Lhs, Rhs expr.Expr
}

func (l LT) GetLhs() expr.Expr {
	return l.Lhs
}

func (l LT) GetRhs() expr.Expr {
	return l.Rhs
}

func (l LT) Eval() interface{} {
	return l.GetLhs().Eval().(int) < l.GetRhs().Eval().(int)
}

func (l LT) Accept(v visit.Visitor) interface{} {
	return v.Visit(l)
}
