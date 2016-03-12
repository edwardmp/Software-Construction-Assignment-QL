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

func (c ComputedQuestion) String() string {
	return fmt.Sprintf("A question with label %s, var decl %s and computation", c.Label, c.VarDecl, c.Computation)
}

func (c ComputedQuestion) GetLabel() interfaces.StrLit {
	return c.Label
}

func (c ComputedQuestion) GetLabelAsString() string {
	return c.Label.GetValue()
}

func (c ComputedQuestion) GetComputation() interfaces.Expr {
	return c.Computation
}

func (c ComputedQuestion) GetVarDecl() interfaces.VarDecl {
	return c.VarDecl
}
