package expr

type StrLit struct {
	Value string
}

func NewStrLit(value string) StrLit {
	return StrLit{Value: value}
}

func (s StrLit) GetValue() string {
	return s.Value
}

func (s StrLit) String() string {
	return s.Value
}
