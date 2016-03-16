package stmt

import (
	"fmt"
	"ql/interfaces"
)

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

func (this ComputedQuestion) String() string {
	return fmt.Sprintf("A question with label %s, var decl %s and computation", this.Label, this.VarDecl, this.Computation)
}

func (this ComputedQuestion) GetLabel() interfaces.StrLit {
	return this.Label
}

func (this ComputedQuestion) GetLabelAsString() string {
	return this.Label.GetValue()
}

func (this ComputedQuestion) GetComputation() interfaces.Expr {
	return this.Computation
}

func (this ComputedQuestion) GetVarDecl() interfaces.VarDecl {
	return this.VarDecl
}
