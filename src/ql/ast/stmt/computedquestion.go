package stmt

import "ql/interfaces"

type ComputedQuestion struct {
	Label       interfaces.StrLit
	VarDecl     interfaces.VarDecl
	Computation interfaces.Expr
	Stmt
}

func NewComputedQuestion(label interfaces.StrLit, varDecl interfaces.VarDecl, computation interfaces.Expr, sourceInfo interface{}) ComputedQuestion {
	return ComputedQuestion{label, varDecl, computation, NewStmt(sourceInfo)}
}

func NewComputedQuestionNoSourceInfo(label interfaces.StrLit, varDecl interfaces.VarDecl, computation interfaces.Expr) ComputedQuestion {
	return NewComputedQuestion(label, varDecl, computation, nil)
}

func (this ComputedQuestion) GetLabel() interfaces.StrLit {
	return this.Label
}

func (this ComputedQuestion) GetComputation() interfaces.Expr {
	return this.Computation
}

func (this ComputedQuestion) GetVarDecl() interfaces.VarDecl {
	return this.VarDecl
}
