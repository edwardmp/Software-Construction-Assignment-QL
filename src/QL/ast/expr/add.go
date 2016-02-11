package expr

type Add struct {
	Lhs, Rhs Expr
}

func (add Add) getLhs() Expr {
	return add.Lhs
}

func (add Add) getRhs() Expr {
	return add.Rhs
}

func (add Add) Eval() interface{} {
	return add.getLhs().Eval().(int) + add.getRhs().Eval().(int)
}
