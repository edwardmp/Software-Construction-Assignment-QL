package expr

type LT struct {
	Lhs, Rhs int
}

func (lt LT) Eval() interface{} {
	return lt.Lhs < lt.Rhs
}
