package expr

type And struct {
	Lhs, Rhs bool
}

func (and And) Eval() interface{} {
	return and.Lhs && and.Rhs
}
