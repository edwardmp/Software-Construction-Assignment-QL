package vari

type VarType interface {
	GetDefaultValue() interface{}
}

type IntType struct {
	VarType
}

func (i IntType) GetDefaultValue() interface{} {
	return 0
}

type StringType struct {
	VarType
}

func (s StringType) GetDefaultValue() interface{} {
	return ""
}

type BoolType struct {
	VarType
}

func (b BoolType) GetDefaultValue() interface{} {
	return false
}
