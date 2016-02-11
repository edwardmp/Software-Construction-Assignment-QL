package expr

type Not struct {
	Value Expr
}

func (not Not) getValue() Expr {
	return not.Value
}

func (not Not) Eval() interface{} {
	return !not.getValue().Eval().(bool)
}
