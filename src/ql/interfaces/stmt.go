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
	GetComputation() Expr
}

type Conditional interface {
	Stmt
	EvalCondition() bool
}

type If interface {
	Conditional
}

type IfElse interface {
	Conditional
}

type StmtList interface {
	Stmt
}
