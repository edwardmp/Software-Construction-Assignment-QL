package stmt

import (
	"fmt"
	"ql/interfaces"
)

type InputQuestion struct {
	Label   interfaces.StrLit
	VarDecl interfaces.VarDecl
	Stmt
}

func NewInputQuestion(label interfaces.StrLit, varDecl interfaces.VarDecl, sourceInfo interface{}) InputQuestion {
	return InputQuestion{label, varDecl, NewStmt(sourceInfo)}
}

func NewInputQuestionNoSourceInfo(label interfaces.StrLit, varDecl interfaces.VarDecl) InputQuestion {
	return NewInputQuestion(label, varDecl, nil)
}

func (i InputQuestion) String() string {
	return fmt.Sprintf("A question with label %s, var decl %s", i.Label, i.VarDecl)
}

func (i InputQuestion) GetLabel() interfaces.StrLit {
	return i.Label
}

func (i InputQuestion) GetLabelAsString() string {
	return i.Label.GetValue()
}

func (i InputQuestion) GetVarDecl() interfaces.VarDecl {
	return i.VarDecl
}
