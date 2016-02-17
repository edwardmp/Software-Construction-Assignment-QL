package binaryoperatorexpr

import ("ql/ast/visit"
"ql/ast/expr"
)

type And struct {
	Lhs, Rhs expr.Expr
}

func (a And) GetLhs() expr.Expr {
	return a.Lhs
}

func (a And) GetRhs() expr.Expr {
	return a.Rhs
}

func (a And) Eval() interface{} {
	return a.GetLhs().(expr.Expr).Eval().(bool) && a.GetRhs().(expr.Expr).Eval().(bool)
}

func (a And) Accept(v visit.Visitor) interface{} {
	return v.Visit(a)
}
