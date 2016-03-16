package interfaces

type Stmt interface {
	ASTNode
	String() string
}

type Form interface {
	Stmt
	GetQuestions() []Question
	GetIdentifier() VarId
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
	EvalCondition(Symbols) bool
	GetCondition() Expr
}

type If interface {
	Conditional
	GetBody() StmtList
}

type IfElse interface {
	Conditional
	GetIfBody() StmtList
	GetElseBody() StmtList
}

type StmtList interface {
	Stmt
	GetConditionals() []Conditional
	GetQuestions() []Question
}
