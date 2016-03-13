package interfaces

type Stmt interface {
	ASTNode
	String() string
}

type Form interface {
	Stmt
	GetQuestions() []Question
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
	GetBody() StmtList
}

type IfElse interface {
	Conditional
}

type StmtList interface {
	Stmt
	GetConditionals() []Conditional
	GetQuestions() []Question
}
