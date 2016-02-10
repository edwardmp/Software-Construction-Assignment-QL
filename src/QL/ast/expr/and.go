package expr

type And struct {
	Lhs, Rhs Expr
}

func (and And) Eval() interface{} {
	return and.Lhs.Eval().(bool) && and.Rhs.Eval().(bool)
}
