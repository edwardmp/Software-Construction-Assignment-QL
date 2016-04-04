package vari

import "ql/interfaces"

type VarDecl struct {
	variableIdentifier interfaces.VarId
	valueType          interfaces.ValueType
	Vari
}

func NewVarDecl(variableIdentifier interfaces.VarId, valueType interfaces.ValueType) VarDecl {
	return VarDecl{variableIdentifier: variableIdentifier, valueType: valueType, Vari: NewVari()}
}

func (this VarDecl) VariableIdentifier() interfaces.VarId {
	return this.variableIdentifier
}

func (this VarDecl) Type() interfaces.ValueType {
	return this.valueType
}
