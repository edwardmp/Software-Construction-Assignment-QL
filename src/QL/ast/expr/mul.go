package expr

type Mul struct {
	Lhs, Rhs Expr
}

func (mul Mul) Eval() interface{} {
	return mul.Lhs.Eval().(int) * mul.Rhs.Eval().(int)
}
