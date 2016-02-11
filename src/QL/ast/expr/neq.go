package expr

type NEq struct {
	Lhs, Rhs Expr
}

func (neq NEq) getLhs() Expr {
	return neq.Lhs
}

func (neq NEq) getRhs() Expr {
	return neq.Rhs
}

func (neq NEq) Eval() interface{} {
	switch neq.Lhs.Eval().(type) {
	case int:
		return neq.getLhs().Eval().(int) != neq.getRhs().Eval().(int)
	case bool:
		return neq.getLhs().Eval().(bool) != neq.getRhs().Eval().(bool)
	default:
		panic("NEq error: comparing unknown types")
	}
}
