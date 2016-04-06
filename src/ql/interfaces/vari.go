package interfaces

type Vari interface {
	ASTNode
}

type VarDecl interface {
	Vari
	VariableIdentifier() VarID
	Type() ValueType
	TypeCheck(TypeCheckArgs)
}

type VarID interface {
	Vari
	Identifier() string
}
