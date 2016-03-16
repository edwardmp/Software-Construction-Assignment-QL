package stmt

import (
	"fmt"
	"ql/interfaces"
)

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

func (f Form) String() string {
	return fmt.Sprintf("A form with identifier %s, statement list %v\n", f.Identifier, f.Content)
}
