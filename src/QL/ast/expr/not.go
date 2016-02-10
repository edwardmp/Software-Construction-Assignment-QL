package expr

type Not struct {
	Value Expr
}

func (not Not) Eval() interface{} {
	return !not.Eval().(bool)
}
