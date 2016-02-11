package expr

type LT struct {
	Lhs, Rhs Expr
}

func (lt LT) getLhs() Expr {
	return lt.Lhs
}

func (lt LT) getRhs() Expr {
	return lt.Rhs
}

func (lt LT) Eval() interface{} {
	return lt.getLhs().Eval().(int) < lt.getRhs().Eval().(int)
}
