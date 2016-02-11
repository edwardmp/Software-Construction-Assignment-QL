package expr

import "ql/ast/visit"

type Eq struct {
	Lhs, Rhs Expr
}

func (e Eq) getLhs() Expr {
	return e.Lhs
}

func (e Eq) getRhs() Expr {
	return e.Rhs
}

func (e Eq) Eval() interface{} {
	switch e.Lhs.Eval().(type) {
	case int:
		return e.getLhs().Eval().(int) == e.getRhs().Eval().(int)
	case bool:
		return e.getLhs().Eval().(bool) == e.getRhs().Eval().(bool)
	default:
		panic("Eq error: comparing unknown types")
	}
}

func (e Eq) Accept(v visit.Visitor) interface{} {
	return v.Visit(e)
}
