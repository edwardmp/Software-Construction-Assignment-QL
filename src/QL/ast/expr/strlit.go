package expr

import "ql/ast/visit"

type StrLit struct {
	Value string
}

func (s StrLit) Eval() interface{} {
	return string(s.Value)
}

func (s StrLit) Accept(v visit.Visitor) interface{} {
	return v.Visit(s)
}
