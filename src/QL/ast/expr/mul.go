package expr

import "ql/ast/visit"

type Mul struct {
	Lhs, Rhs Expr
}

func (m Mul) getLhs() Expr {
	return m.Lhs
}

func (m Mul) getRhs() Expr {
	return m.Rhs
}

func (m Mul) Eval() interface{} {
	return m.getLhs().Eval().(int) * m.getRhs().Eval().(int)
}

func (m Mul) Accept(v visit.Visitor) interface{} {
	return v.Visit(m)
}
