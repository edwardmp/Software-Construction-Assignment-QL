package vari

import (
	"ql/ast/visit"
)

type VarDecl struct {
	Ident VarId
	Type  variType
}

// TODO: needed to break import cycle, look at better solution
type variType interface {
	GetDefaultValue() interface{}
}

func (va VarDecl) GetType() variType {
	return va.Type
}

func (va VarDecl) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(va, s)
}
