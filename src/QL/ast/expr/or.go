package expr

type Or struct {
	Lhs, Rhs bool
}

func (or Or) Eval() interface{} {
	return or.Lhs || or.Rhs
}
