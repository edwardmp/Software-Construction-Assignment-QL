package stmt

import "ql/interfaces"

type Form struct {
	Identifier interfaces.VarId
	Content    StmtList
	Stmt
}

func NewForm(identifier interfaces.VarId, content StmtList, sourceInfo interface{}) Form {
	return Form{identifier, content, NewStmt(sourceInfo)}
}

func NewFormNoSourceInfo(identifier interfaces.VarId, content StmtList) Form {
	return NewForm(identifier, content, nil)
}

func (this Form) GetQuestions() []interfaces.Question {
	return this.Content.Questions
}

func (this Form) GetIdentifier() interfaces.VarId {
	return this.Identifier
}
