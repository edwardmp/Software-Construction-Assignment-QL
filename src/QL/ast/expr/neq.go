package expr

type NEq struct {
	Lhs, Rhs interface{}
}

func (neq NEq) Eval() interface{} {
	return neq.Lhs != neq.Rhs
}
