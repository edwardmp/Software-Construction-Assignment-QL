package stmt

import "ql/interfaces"

type Form struct {
	Identifier interfaces.VarId
	Content    StmtList
	Stmt
}

func NewForm(identifier interfaces.VarId, content StmtList) Form {
	return Form{identifier, content, NewStmt()}
}

func (this Form) GetQuestions() []interfaces.Question {
	return this.Content.Questions
}

func (this Form) GetIdentifier() interfaces.VarId {
	return this.Identifier
}
