package vari

import (
	"ql/ast/visit"
)

type VarDecl struct {
	Ident VarId
	Type  VarTypeId
}

func (va VarDecl) GetType() VarTypeId {
	return va.Type
}

func (va VarDecl) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(va, s)
}
