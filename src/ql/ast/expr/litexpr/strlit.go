package litexpr

type StrLit struct {
	Value string
}

func (s StrLit) Eval(sy interface{}) interface{} {
	return string(s.Value)
}

func (s StrLit) String() string {
	return s.Value
}
