package expr

type GT struct {
	Lhs, Rhs Expr
}

func (gt GT) getLhs() Expr {
	return gt.Lhs
}

func (gt GT) getRhs() Expr {
	return gt.Rhs
}

func (gt GT) Eval() interface{} {
	return gt.getLhs().Eval().(int) > gt.getRhs().Eval().(int)
}
