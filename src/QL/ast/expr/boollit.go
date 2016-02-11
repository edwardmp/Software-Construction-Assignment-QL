package expr

import "ql/ast/visit"

type BoolLit struct {
	Value bool
}

func (b BoolLit) Eval() interface{} {
	return bool(b.Value)
}

func (b BoolLit) Accept(v visit.Visitor) interface{} {
	return v.Visit(b)
}
