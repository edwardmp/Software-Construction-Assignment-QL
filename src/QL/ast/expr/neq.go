package expr

import "ql/ast/visit"

type NEq struct {
	Lhs, Rhs Expr
}

func (n NEq) getLhs() Expr {
	return n.Lhs
}

func (n NEq) getRhs() Expr {
	return n.Rhs
}

func (n NEq) Eval() interface{} {
	switch n.Lhs.Eval().(type) {
	case int:
		return n.getLhs().Eval().(int) != n.getRhs().Eval().(int)
	case bool:
		return n.getLhs().Eval().(bool) != n.getRhs().Eval().(bool)
	default:
		panic("NEq error: comparing unknown types")
	}
}

func (n NEq) Accept(v visit.Visitor) interface{} {
	return v.Visit(n)
}
