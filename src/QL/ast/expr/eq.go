package expr

type Eq struct {
	Lhs, Rhs interface{}
}

func (eq Eq) Eval() interface{} {
	return eq.Lhs == eq.Rhs
}
