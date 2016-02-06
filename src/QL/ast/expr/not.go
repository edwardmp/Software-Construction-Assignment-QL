package expr

type Not struct {
	Value bool
}

func (not Not) Eval() interface{} {
	return !not.Value
}
