package vari

import "ql/interfaces"

type VarDecl struct {
	Ident interfaces.VarId
	Type  interfaces.VarType
}

func (v VarDecl) GetIdent() interfaces.VarId {
	return v.Ident
}

func (v VarDecl) GetType() interfaces.VarType {
	return v.Type
}
