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

func (this VarDecl) GetIdent() interfaces.VarId {
	return this.Ident
}

func (this VarDecl) GetType() interfaces.VarType {
	return this.Type
}
