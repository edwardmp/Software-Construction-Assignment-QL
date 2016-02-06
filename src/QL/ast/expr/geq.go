package expr

type GEq struct {
	Lhs, Rhs int
}

func (geq GEq) Eval() interface{} {
	return geq.Lhs >= geq.Rhs
}
