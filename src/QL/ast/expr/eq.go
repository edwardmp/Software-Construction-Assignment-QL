package expr

import "ql/ast/visit"

type Eq struct {
	Lhs, Rhs Expr
}

func (e Eq) GetLhs() Expr {
	return e.Lhs
}

func (e Eq) GetRhs() Expr {
	return e.Rhs
}

func (e Eq) Eval() interface{} {
	switch e.Lhs.Eval().(type) {
	case int:
		return e.GetLhs().Eval().(int) == e.GetRhs().Eval().(int)
	case bool:
		return e.GetLhs().Eval().(bool) == e.GetRhs().Eval().(bool)
	default:
		panic("Eq error: comparing unknown types")
	}
}

func (e Eq) Accept(v visit.Visitor) interface{} {
	return v.Visit(e)
}
