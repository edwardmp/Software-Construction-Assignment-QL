package interfaces

type VarDecl interface {
	Node
	GetIdent() VarId
	GetType() VarType
}

type VarId interface {
	Node
	GetIdent() string
}

type VarType interface {
	Node
	GetDefaultValue() interface{}
}

type IntType interface {
	VarType
}

type BoolType interface {
	VarType
}

type StringType interface {
	VarType
}
