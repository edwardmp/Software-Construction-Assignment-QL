package interfaces

type Var interface {
	ASTNode
}

type VarDecl interface {
	Var
	GetIdent() VarId
	GetType() ValueType
	TypeCheck(TypeChecker, TypeCheckSymbols) interface{}
}

type VarId interface {
	Var
	GetIdent() string
}
