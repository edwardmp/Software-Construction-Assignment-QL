package expr

import "ql/ast/visit"

type Mul struct {
	Lhs, Rhs Expr
}

func (m Mul) GetLhs() Expr {
	return m.Lhs
}

func (m Mul) GetRhs() Expr {
	return m.Rhs
}

func (m Mul) Eval() interface{} {
	return m.GetLhs().Eval().(int) * m.GetRhs().Eval().(int)
}

func (m Mul) Accept(v visit.Visitor) interface{} {
	return v.Visit(m)
}
