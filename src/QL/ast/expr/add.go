package expr

type Add struct {
	Lhs, Rhs int
}

func (add Add) Eval() interface{} {
	return add.Lhs + add.Rhs
}
