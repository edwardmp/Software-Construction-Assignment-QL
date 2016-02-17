package vari

import "ql/ast/visit"

type VarId struct {
	Ident string
}

func (va VarId) Accept(v visit.Visitor) interface{} {
	return v.Visit(va)
}
