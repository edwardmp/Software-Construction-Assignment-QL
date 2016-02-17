package binaryoperatorexpr

import ("ql/ast/visit"
"ql/ast/expr")

type Mul struct {
	Lhs, Rhs expr.Expr
}

func (m Mul) GetLhs() expr.Expr {
	return m.Lhs
}

func (m Mul) GetRhs() expr.Expr {
	return m.Rhs
}

func (m Mul) Eval() interface{} {
	return m.GetLhs().Eval().(int) * m.GetRhs().Eval().(int)
}

func (m Mul) Accept(v visit.Visitor) interface{} {
	return v.Visit(m)
}
