package vari

import "ql/interfaces"

type VarDecl struct {
	Ident interfaces.VarId
	Type  interfaces.VarType
	Var
}

func NewVarDecl(varIdent interfaces.VarId, varType interfaces.VarType, sourceInfo interface{}) VarDecl {
	return VarDecl{varIdent, varType, NewVar(sourceInfo)}
}

func NewVarDeclNoSourceInfo(varIdent interfaces.VarId, varType interfaces.VarType) VarDecl {
	return NewVarDecl(varIdent, varType, nil)
}

func (v VarDecl) GetIdent() interfaces.VarId {
	return v.Ident
}

func (v VarDecl) GetType() interfaces.VarType {
	return v.Type
}
