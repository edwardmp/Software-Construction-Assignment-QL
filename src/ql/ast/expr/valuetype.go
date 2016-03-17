package expr

import "ql/interfaces"

type ValueType struct {
	Expr
	TypeString string
}

func NewValueType(typeString string, sourceInfo interface{}) ValueType {
	return ValueType{NewExpr(sourceInfo), typeString}
}

func (this ValueType) String() string {
	return this.TypeString
}

/* UnknownType */

type UnknownType struct {
	ValueType
}

func NewUnknownType() UnknownType {
	return UnknownType{NewValueType("Unknown", nil)}
}

func (this UnknownType) GetDefaultValue() interfaces.LitExpr {
	panic("UnknownType has no default value")

	return nil
}

/* IntType */

type IntType struct {
	ValueType
}

func NewIntType(sourceInfo interface{}) IntType {
	return IntType{NewValueType("Integer", sourceInfo)}
}

func NewIntTypeNoSourceInfo() IntType {
	return NewIntType(nil)
}

func (this IntType) GetDefaultValue() interfaces.LitExpr {
	return NewIntLitNoSourceInfo(0)
}

/* StringType */

type StringType struct {
	ValueType
}

func NewStringType(sourceInfo interface{}) StringType {
	return StringType{NewValueType("String", sourceInfo)}
}

func NewStringTypeNoSourceInfo() StringType {
	return NewStringType(nil)
}

func (this StringType) GetDefaultValue() interfaces.LitExpr {
	return NewStrLitNoSourceInfo("")
}

func (this StringType) String() string {
	return this.TypeString
}

/* BoolType */

type BoolType struct {
	ValueType
}

func NewBoolType(sourceInfo interface{}) BoolType {
	return BoolType{NewValueType("Boolean", sourceInfo)}
}

func NewBoolTypeNoSourceInfo() BoolType {
	return NewBoolType(nil)
}

func (this BoolType) GetDefaultValue() interfaces.LitExpr {
	return NewBoolLitNoSourceInfo(false)
}
