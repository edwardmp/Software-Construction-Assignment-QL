package vari

import "ql/interfaces"

type VarDecl struct {
	variableIdentifier interfaces.VarId
	valueType          interfaces.ValueType
	Vari
}

func NewVarDecl(varIdent interfaces.VarId, varType interfaces.ValueType) VarDecl {
	return VarDecl{varIdent, varType, NewVari()}
}

func (this VarDecl) VariableIdentifier() interfaces.VarId {
	return this.variableIdentifier
}

func (this VarDecl) Type() interfaces.ValueType {
	return this.valueType
}
