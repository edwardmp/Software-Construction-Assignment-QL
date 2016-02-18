package lit

import (
	"ql/ast/visit"
)

type StrLit struct {
	Value string
}

func (s StrLit) GetValue() interface{} {
	return s.Value
}

func (s StrLit) Eval() interface{} {
	return string(s.Value)
}

func (s StrLit) Accept(v visit.Visitor, sy interface{}) interface{} {
	return v.Visit(s, sy)
}
