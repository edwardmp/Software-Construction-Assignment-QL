package vari

import (
	"ql/ast/expr"
)

type IntType struct {
}

func (i IntType) GetDefaultValue() interface{} {
	return expr.NewIntLit(0)
}

type StringType struct {
}

func (s StringType) GetDefaultValue() interface{} {
	return expr.NewStrLit("")
}

type BoolType struct {
}

func (b BoolType) GetDefaultValue() interface{} {
	return expr.NewBoolLit(false)
}
