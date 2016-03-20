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

func NewStmtList(questions []interfaces.Question, conditionals []interfaces.Conditional) StmtList {
	return StmtList{questions, conditionals, NewStmt()}
}

func NewEmptyStmtList() StmtList {
	return NewStmtList(nil, nil)
}

func (this StmtList) GetQuestions() []interfaces.Question {
	return this.Questions
}

func (this StmtList) GetConditionals() []interfaces.Conditional {
	return this.Conditionals
}

func (this StmtList) AddToCorrectSlice(questionOrConditional interface{}) StmtList {
	switch assertedStmt := questionOrConditional.(type) {
	default:
		panic(fmt.Errorf("Unexpected StmtList type passed to AddToCorrectSlice %T\n", assertedStmt))
	case interfaces.Question:
		this.Questions = append(this.Questions, assertedStmt)
	case interfaces.Conditional:
		this.Conditionals = append(this.Conditionals, assertedStmt)
	}

	return this
}
