package interfaces

type Vari interface {
	ASTNode
}

type VarDecl interface {
	Vari
	VariableIdentifier() VarId
	Type() ValueType
	TypeCheck(TypeCheckArgs)
}

type VarId interface {
	Vari
	Identifier() string
}
