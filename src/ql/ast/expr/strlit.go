package expr

import "ql/interfaces"

type StrLit struct {
	Value string
	interfaces.Expr
}

func NewStrLit(value string) StrLit {
	return StrLit{Value: value}
}

func (s StrLit) GetValue() string {
	return s.Value
}

func (s StrLit) Eval(sy interface{}) interface{} {
	return string(s.Value)
}

func (s StrLit) String() string {
	return s.Value
}
