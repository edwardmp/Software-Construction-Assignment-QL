package stmt

import (
	"fmt"
	"ql/interfaces"
)

type StmtList struct {
	Questions    []interfaces.Question
	Conditionals []interfaces.Conditional
	Stmt
}

func NewStmtList(questions []interfaces.Question, conditionals []interfaces.Conditional, sourceInfo interface{}) StmtList {
	return StmtList{questions, conditionals, NewStmt(sourceInfo)}
}

func NewStmtListNoSourceInfo(questions []interfaces.Question, conditionals []interfaces.Conditional) StmtList {
	return NewStmtList(questions, conditionals, nil)
}

func NewEmptyStmtList(sourceInfo interface{}) StmtList {
	return StmtList{Stmt: NewStmt(sourceInfo)}
}

func NewEmptyStmtListNoSourceInfo() StmtList {
	return NewEmptyStmtList(nil)
}

func (this StmtList) GetQuestions() []interfaces.Question {
	return this.Questions
}

func (this StmtList) GetConditionals() []interfaces.Conditional {
	return this.Conditionals
}

func (this StmtList) AddToCorrectSlice(i interface{}) StmtList {
	switch t := i.(type) {
	default:
		panic(fmt.Sprintf("Unexpected StmtList type passed %T\n", t))
	case interfaces.Question:
		this.Questions = append(this.Questions, i.(interfaces.Question))
	case If:
		this.Conditionals = append(this.Conditionals, i.(If))
	case IfElse:
		this.Conditionals = append(this.Conditionals, i.(IfElse))
	}

	return this
}

func (this StmtList) String() string {
	return fmt.Sprintf("A statement list with %d questions and %d conditionals", len(this.Questions), len(this.Conditionals))
}
