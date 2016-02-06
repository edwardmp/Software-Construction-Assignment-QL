package expr

type Div struct {
	Lhs, Rhs int
}

func (div Div) Eval() interface{} {
	return div.Lhs / div.Rhs
}
