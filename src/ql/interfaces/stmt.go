package interfaces

type Stmt interface {
	Node
	String() string
}

type Form interface {
	Stmt
}

type Question interface {
	Stmt
	GetLabel() StrLit
	GetLabelAsString() string
	GetVarDecl() VarDecl
}

type InputQuestion interface {
	Question
}

type ComputedQuestion interface {
	Question
}

type If interface {
	Stmt
}

type IfElse interface {
	Stmt
}

type StmtList interface {
	Stmt
}
