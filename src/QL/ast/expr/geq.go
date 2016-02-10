package expr

type GEq struct {
	Lhs, Rhs Expr
}

func (geq GEq) Eval() interface{} {
	return geq.Lhs.Eval().(int) >= geq.Rhs.Eval().(int)
}
