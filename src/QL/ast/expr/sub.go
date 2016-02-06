package expr

type Sub struct {
	Lhs, Rhs int
}

func (sub Sub) Eval() interface{} {
	return sub.Lhs - sub.Rhs
}
