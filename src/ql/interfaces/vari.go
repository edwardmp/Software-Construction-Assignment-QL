package interfaces

type Var interface {
	ASTNode
}

type VarDecl interface {
	Var
	GetIdent() VarId
	GetType() ValueType
	TypeCheck(TypeCheckArgs) interface{}
}

type VarId interface {
	Var
	GetIdent() string
}
