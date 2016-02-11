package expr

type Sub struct {
	Lhs, Rhs Expr
}

func (sub Sub) getLhs() Expr {
	return sub.Lhs
}

func (sub Sub) getRhs() Expr {
	return sub.Rhs
}

func (sub Sub) Eval() interface{} {
	return sub.getLhs().Eval().(int) - sub.getRhs().Eval().(int)
}
