package interfaces

type Var interface {
	ASTNode
}

type VarDecl interface {
	Var
	Identifier() VarId
	Type() ValueType
	TypeCheck(TypeCheckArgs) interface{}
}

type VarId interface {
	Var
	Identifier() string
}
