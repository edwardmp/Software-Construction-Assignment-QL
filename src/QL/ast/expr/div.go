package expr

type Div struct {
	Lhs, Rhs Expr
}

func (div Div) getLhs() Expr {
	return div.Lhs
}

func (div Div) getRhs() Expr {
	return div.Rhs
}

func (div Div) Eval() interface{} {
	return div.getLhs().Eval().(int) / div.getRhs().Eval().(int)
}
