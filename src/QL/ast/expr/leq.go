package expr

type LEq struct {
	Lhs, Rhs Expr
}

func (leq LEq) Eval() interface{} {
	return leq.Lhs.Eval().(int) <= leq.Rhs.Eval().(int)
}
