package expr

type Eq struct {
	Lhs, Rhs Expr
}

func (eq Eq) Eval() interface{} {
	switch eq.Lhs.Eval().(type) {
	case int:
		return eq.Lhs.Eval().(int) == eq.Rhs.Eval().(int)
	case bool:
		return eq.Lhs.Eval().(bool) == eq.Rhs.Eval().(bool)
	default:
		panic("Eq error: comparing unknown types")
	}
}
