package expr

type GT struct {
	Lhs, Rhs int
}

func (gt GT) Eval() interface{} {
	return gt.Lhs > gt.Rhs
}
