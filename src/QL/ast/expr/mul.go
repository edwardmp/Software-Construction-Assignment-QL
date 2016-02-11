package expr

type Mul struct {
	Lhs, Rhs Expr
}

func (mul Mul) getLhs() Expr {
	return mul.Lhs
}

func (mul Mul) getRhs() Expr {
	return mul.Rhs
}

func (mul Mul) Eval() interface{} {
	return mul.getLhs().Eval().(int) * mul.getRhs().Eval().(int)
}
