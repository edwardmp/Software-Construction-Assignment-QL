package expr

type Add struct {
	Lhs, Rhs interface{}
}

func (add Add) Eval() interface{} {
	return add.Lhs.(Expr).Eval().(int) + add.Rhs.(Expr).Eval().(int)
}
