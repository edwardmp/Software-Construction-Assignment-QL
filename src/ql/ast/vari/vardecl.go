package vari

import "ql/interfaces"

type VarDecl struct {
	ident     interfaces.VarId
	valueType interfaces.ValueType
	Var
}

func NewVarDecl(varIdent interfaces.VarId, varType interfaces.ValueType) VarDecl {
	return VarDecl{varIdent, varType, NewVar()}
}

func (this VarDecl) Identifier() interfaces.VarId {
	return this.ident
}

func (this VarDecl) Type() interfaces.ValueType {
	return this.valueType
}
