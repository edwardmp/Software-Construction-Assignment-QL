package lit

import (
	"ql/ast/visit"
)

type IntLit struct {
	Value int
}

func (i IntLit) GetValue() interface{} {
	return i.Value
}

func (i IntLit) Eval() interface{} {
	return i.Value
}

func (i IntLit) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(i, s)
}
