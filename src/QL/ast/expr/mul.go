package expr

type Mul struct {
	Lhs, Rhs int
}

func (mul Mul) Eval() interface{} {
	return mul.Lhs * mul.Rhs
}
