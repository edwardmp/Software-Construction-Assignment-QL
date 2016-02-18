package vari

import (
	"ql/ast/visit"
)

type VarId struct {
	Ident string
}

func (va VarId) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(va, s)
}

func (va VarId) String() string {
	return va.Ident
}
