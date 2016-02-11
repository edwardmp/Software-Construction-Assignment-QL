package expr

type Eq struct {
	Lhs, Rhs Expr
}

func (eq Eq) getLhs() Expr {
	return eq.Lhs
}

func (eq Eq) getRhs() Expr {
	return eq.Rhs
}

func (eq Eq) Eval() interface{} {
	switch eq.Lhs.Eval().(type) {
	case int:
		return eq.getLhs().Eval().(int) == eq.getRhs().Eval().(int)
	case bool:
		return eq.getLhs().Eval().(bool) == eq.getRhs().Eval().(bool)
	default:
		panic("Eq error: comparing unknown types")
	}
}
