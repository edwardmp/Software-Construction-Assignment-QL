package expr

type StrLit struct {
	Value string
}

func (s StrLit) Eval() interface{} {
	return string(s.Value)
}
