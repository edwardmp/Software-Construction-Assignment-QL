package interfaces

type VarDecl interface {
	ASTNode
	GetIdent() VarId
	GetType() VarType
}

type VarId interface {
	ASTNode
	GetIdent() string
}

type VarType interface {
	ASTNode
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
