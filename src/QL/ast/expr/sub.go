package expr

type Sub struct {
	Lhs, Rhs Expr
}

func (sub Sub) Eval() interface{} {
	return sub.Lhs.Eval().(int) - sub.Rhs.Eval().(int)
}
