package stmt

import "ql/interfaces"

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

func (this InputQuestion) GetLabel() interfaces.StrLit {
	return this.Label
}

func (this InputQuestion) GetLabelAsString() string {
	return this.Label.GetValue()
}

func (this InputQuestion) GetVarDecl() interfaces.VarDecl {
	return this.VarDecl
}
