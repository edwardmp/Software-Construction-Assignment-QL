package expr

type LEq struct {
	Lhs, Rhs int
}

func (leq LEq) Eval() interface{} {
	return leq.Lhs <= leq.Rhs
}
