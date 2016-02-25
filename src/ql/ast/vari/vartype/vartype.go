package vartype

import "ql/ast/expr/lit"

type VarType interface {
	GetDefaultValue() interface{}
}

type IntType struct {
	VarType
}

func (i IntType) GetDefaultValue() interface{} {
	return lit.IntLit{0}
}

type StringType struct {
	VarType
}

func (s StringType) GetDefaultValue() interface{} {
	return lit.StrLit{""}
}

type BoolType struct {
	VarType
}

func (b BoolType) GetDefaultValue() interface{} {
	return lit.BoolLit{false}
}
