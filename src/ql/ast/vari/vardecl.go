package vari

import "ql/interfaces"

type VarDecl struct {
	Ident interfaces.VarId
	Type  interfaces.ValueType
	Var
}

func NewVarDecl(varIdent interfaces.VarId, varType interfaces.ValueType, sourceInfo interface{}) VarDecl {
	return VarDecl{varIdent, varType, NewVar(sourceInfo)}
}

func NewVarDeclNoSourceInfo(varIdent interfaces.VarId, varType interfaces.ValueType) VarDecl {
	return NewVarDecl(varIdent, varType, nil)
}

func (this VarDecl) GetIdent() interfaces.VarId {
	return this.Ident
}

func (this VarDecl) GetType() interfaces.ValueType {
	return this.Type
}
