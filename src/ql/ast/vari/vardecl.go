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

func (v VarDecl) GetType() variType {
	return v.Type
}

func (v VarDecl) GetIdentifier() VarId {
	return v.Ident
}

func (va VarDecl) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(va, s)
}
