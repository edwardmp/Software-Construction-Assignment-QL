package expr

type Or struct {
	Lhs, Rhs Expr
}

func (or Or) getLhs() Expr {
	return or.Lhs
}

func (or Or) getRhs() Expr {
	return or.Rhs
}

func (or Or) Eval() interface{} {
	return or.getLhs().Eval().(bool) || or.getRhs().Eval().(bool)
}
