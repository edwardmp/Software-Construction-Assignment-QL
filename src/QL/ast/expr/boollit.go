package expr

type BoolLit struct {
	Value bool
}

func (i BoolLit) Eval() interface{} {
	return bool(i.Value)
}
