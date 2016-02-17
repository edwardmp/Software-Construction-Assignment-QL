package binaryoperatorexpr

import ("ql/ast/visit"
"ql/ast/expr")

type LEq struct {
	Lhs, Rhs expr.Expr
}

func (l LEq) GetLhs() expr.Expr {
	return l.Lhs
}

func (l LEq) GetRhs() expr.Expr {
	return l.Rhs
}

func (l LEq) Eval() interface{} {
	return l.GetLhs().Eval().(int) <= l.GetRhs().Eval().(int)
}

func (l LEq) Accept(v visit.Visitor) interface{} {
	return v.Visit(l)
}