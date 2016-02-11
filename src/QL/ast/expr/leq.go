package expr

type LEq struct {
	Lhs, Rhs Expr
}

func (leq LEq) getLhs() Expr {
	return leq.Lhs
}

func (leq LEq) getRhs() Expr {
	return leq.Rhs
}

func (leq LEq) Eval() interface{} {
	return leq.getLhs().Eval().(int) <= leq.getRhs().Eval().(int)
}
