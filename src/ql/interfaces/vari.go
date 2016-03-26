package interfaces

type Vari interface {
	ASTNode
}

type VarDecl interface {
	Vari
	Identifier() VarId
	Type() ValueType
	TypeCheck(TypeCheckArgs) interface{}
}

type VarId interface {
	Vari
	Identifier() string
}
