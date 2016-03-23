package expr

import "ql/interfaces"

type ValueType struct {
	Expr
	TypeString string
}

func NewValueType(typeString string) ValueType {
	return ValueType{NewExpr(), typeString}
}

func (this ValueType) String() string {
	return this.TypeString
}

/* UnknownType */

type UnknownType struct {
	ValueType
}

func NewUnknownType() UnknownType {
	return UnknownType{NewValueType("Unknown")}
}

func (this UnknownType) DefaultValue() interfaces.LitExpr {
	panic("UnknownType has no default value")

	return nil
}

/* IntType */

type IntType struct {
	ValueType
}

func NewIntType() IntType {
	return IntType{NewValueType("Integer")}
}

func (this IntType) DefaultValue() interfaces.LitExpr {
	return NewIntLit(0)
}

/* StringType */

type StringType struct {
	ValueType
}

func NewStringType() StringType {
	return StringType{NewValueType("String")}
}

func (this StringType) DefaultValue() interfaces.LitExpr {
	return NewStrLit("")
}

func (this StringType) String() string {
	return this.TypeString
}

/* BoolType */

type BoolType struct {
	ValueType
}

func NewBoolType() BoolType {
	return BoolType{NewValueType("Boolean")}
}

func (this BoolType) DefaultValue() interfaces.LitExpr {
	return NewBoolLit(false)
}
