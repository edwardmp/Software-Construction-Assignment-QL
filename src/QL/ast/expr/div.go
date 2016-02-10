package expr

type Div struct {
	Lhs, Rhs Expr
}

func (div Div) Eval() interface{} {
	return div.Lhs.Eval().(int) / div.Rhs.Eval().(int)
}
