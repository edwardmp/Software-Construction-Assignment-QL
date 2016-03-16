package expr

type ValueType struct {
	Expr
}

func NewValueType(sourceInfo interface{}) ValueType {
	return ValueType{NewExpr(sourceInfo)}
}

type IntType struct {
	ValueType
	TypeString string
}

func NewIntType(sourceInfo interface{}) IntType {
	return IntType{NewValueType(sourceInfo), "Integer"}
}

func NewIntTypeNoSourceInfo() IntType {
	return NewIntType(nil)
}

func (this IntType) GetDefaultValue() interface{} {
	return NewIntLitNoSourceInfo(0)
}

func (this IntType) String() string {
	return this.TypeString
}

type StringType struct {
	ValueType
	TypeString string
}

func NewStringType(sourceInfo interface{}) StringType {
	return StringType{NewValueType(sourceInfo), "String"}
}

func NewStringTypeNoSourceInfo() StringType {
	return NewStringType(nil)
}

func (this StringType) GetDefaultValue() interface{} {
	return NewStrLitNoSourceInfo("")
}

func (this StringType) String() string {
	return this.TypeString
}

type BoolType struct {
	ValueType
	TypeString string
}

func NewBoolType(sourceInfo interface{}) BoolType {
	return BoolType{NewValueType(sourceInfo), "Boolean"}
}

func NewBoolTypeNoSourceInfo() BoolType {
	return NewBoolType(nil)
}

func (this BoolType) GetDefaultValue() interface{} {
	return NewBoolLitNoSourceInfo(false)
}

func (this BoolType) String() string {
	return this.TypeString
}
