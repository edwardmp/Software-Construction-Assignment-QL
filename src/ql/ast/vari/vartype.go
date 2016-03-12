package vari

import (
	"ql/ast/expr"
)

type VarType struct {
	Var
}

func NewVarType(sourceInfo interface{}) VarType {
	return VarType{NewVar(sourceInfo)}
}

type IntType struct {
	VarType
}

func NewIntType(sourceInfo interface{}) IntType {
	return IntType{NewVarType(sourceInfo)}
}

func NewIntTypeNoSourceInfo() IntType {
	return NewIntType(nil)
}

func (i IntType) GetDefaultValue() interface{} {
	return expr.NewIntLitNoSourceInfo(0)
}

func NewStringType(sourceInfo interface{}) StringType {
	return StringType{NewVarType(sourceInfo)}
}

func NewStringTypeNoSourceInfo() StringType {
	return StringType{NewVarType(nil)}
}

type StringType struct {
	VarType
}

func (s StringType) GetDefaultValue() interface{} {
	return expr.NewStrLitNoSourceInfo("")
}

func NewBoolType(sourceInfo interface{}) BoolType {
	return BoolType{NewVarType(sourceInfo)}
}

func NewBoolTypeNoSourceInfo() BoolType {
	return BoolType{NewVarType(nil)}
}

type BoolType struct {
	VarType
}

func (b BoolType) GetDefaultValue() interface{} {
	return expr.NewBoolLitNoSourceInfo(false)
}
