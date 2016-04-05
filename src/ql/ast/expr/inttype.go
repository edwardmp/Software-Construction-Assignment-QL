package expr

import "ql/interfaces"

type IntType struct {
	ValueType
}

func NewIntType() IntType {
	return IntType{NewValueType("Integer")}
}

func (this IntType) DefaultValue() interfaces.LitExpr {
	return NewIntegerLiteral(0)
}
