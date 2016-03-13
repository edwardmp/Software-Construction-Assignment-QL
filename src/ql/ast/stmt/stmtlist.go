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

func (s StmtList) AddToCorrectSlice(i interface{}) StmtList {
	switch t := i.(type) {
	default:
		panic(fmt.Sprintf("Unexpected StmtList type passed %T\n", t))
	case interfaces.Question:
		s.Questions = append(s.Questions, i.(interfaces.Question))
	case If:
		s.Conditionals = append(s.Conditionals, i.(If))
	case IfElse:
		s.Conditionals = append(s.Conditionals, i.(IfElse))
	}

	return s
}

func (s StmtList) String() string {
	return fmt.Sprintf("A statement list with %d questions and %d conditionals", len(s.Questions), len(s.Conditionals))
}
