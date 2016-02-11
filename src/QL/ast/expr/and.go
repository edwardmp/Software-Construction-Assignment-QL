package expr

type And struct {
	Lhs, Rhs Expr
}

func (and And) getLhs() Expr {
	return and.Lhs
}

func (and And) getRhs() Expr {
	return and.Rhs
}

func (and And) Eval() interface{} {
	return and.getLhs().(Expr).Eval().(bool) && and.getRhs().(Expr).Eval().(bool)
}
