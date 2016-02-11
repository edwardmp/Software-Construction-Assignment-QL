package expr

type GEq struct {
	Lhs, Rhs Expr
}

func (geq GEq) getLhs() Expr {
	return geq.Lhs
}

func (geq GEq) getRhs() Expr {
	return geq.Rhs
}

func (geq GEq) Eval() interface{} {
	return geq.getLhs().Eval().(int) >= geq.getRhs().Eval().(int)
}
