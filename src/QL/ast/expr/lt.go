package expr

type LT struct {
	Lhs, Rhs Expr
}

func (lt LT) Eval() interface{} {
	return lt.Lhs.Eval().(int) < lt.Rhs.Eval().(int)
}
