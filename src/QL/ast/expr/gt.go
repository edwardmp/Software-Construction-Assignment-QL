package expr

type GT struct {
	Lhs, Rhs Expr
}

func (gt GT) Eval() interface{} {
	return gt.Lhs.Eval().(int) > gt.Rhs.Eval().(int)
}
