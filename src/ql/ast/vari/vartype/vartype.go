package vartype

import "ql/ast/expr/litexpr"

type VarType interface {
	GetDefaultValue() interface{}
}

type IntType struct {
	VarType
}

func (i IntType) GetDefaultValue() interface{} {
	return litexpr.IntLit{0}
}

type StringType struct {
	VarType
}

func (s StringType) GetDefaultValue() interface{} {
	return litexpr.StrLit{""}
}

type BoolType struct {
	VarType
}

func (b BoolType) GetDefaultValue() interface{} {
	return litexpr.BoolLit{false}
}
