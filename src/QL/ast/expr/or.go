package expr

type Or struct {
	Lhs, Rhs Expr
}

func (or Or) Eval() interface{} {
	return or.Lhs.Eval().(bool) || or.Rhs.Eval().(bool)
}
