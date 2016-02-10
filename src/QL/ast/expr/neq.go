package expr

type NEq struct {
	Lhs, Rhs Expr
}

func (neq NEq) Eval() interface{} {
	switch neq.Lhs.Eval().(type) {
	case int:
		return neq.Lhs.Eval().(int) != neq.Rhs.Eval().(int)
	case bool:
		return neq.Lhs.Eval().(bool) != neq.Rhs.Eval().(bool)
	default:
		panic("NEq error: comparing unknown types")
	}
}
