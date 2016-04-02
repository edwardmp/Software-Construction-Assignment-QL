package stmt

import (
	"fmt"
	"ql/interfaces"
)

type StmtList struct {
	questions    []interfaces.Question
	conditionals []interfaces.Conditional
	Stmt
}

func NewStmtList(questions []interfaces.Question, conditionals []interfaces.Conditional) StmtList {
	return StmtList{questions, conditionals, NewStmt()}
}

func NewEmptyStmtList() StmtList {
	return NewStmtList(nil, nil)
}

func (this StmtList) Questions() []interfaces.Question {
	return this.questions
}

func (this StmtList) Conditionals() []interfaces.Conditional {
	return this.conditionals
}

func (this StmtList) AddToCorrectSlice(questionOrConditional interfaces.Stmt) StmtList {
	switch assertedStmt := questionOrConditional.(type) {
	case interfaces.Question:
		this.questions = append(this.questions, assertedStmt)
	case interfaces.Conditional:
		this.conditionals = append(this.conditionals, assertedStmt)
	default:
		panic(fmt.Errorf("Unexpected StmtList type passed to AddToCorrectSlice %T\n", assertedStmt))
	}

	return this
}
