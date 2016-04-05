package interfaces

type Stmt interface {
	ASTNode
	TypeCheck(TypeCheckArgs)
}

type Form interface {
	Stmt
	Content() StmtList
	Identifier() VarId
	Questions() []Question
}

type Question interface {
	Stmt
	Label() StrLit
	VarDecl() VarDecl
}

type InputQuestion interface {
	Question
}

type ComputedQuestion interface {
	Question
	Computation() Expr
}

type Conditional interface {
	Stmt
	EvalCondition(VarIdValueSymbols) BoolValue
	Condition() Expr
}

type If interface {
	Conditional
	Body() StmtList
}

type IfElse interface {
	Conditional
	IfBody() StmtList
	ElseBody() StmtList
}

type StmtList interface {
	Stmt
	Conditionals() []Conditional
	Questions() []Question
}
