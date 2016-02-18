package stmt

import (
	"fmt"
	"ql/ast/visit"
)

type StmtList struct {
	Questions    []Question
	Conditionals []Conditional
}

func (s StmtList) Eval() {
	for _, conditional := range s.Conditionals {
		conditional.Eval()
	}
}

func (s StmtList) AddToCorrectSlice(i interface{}) StmtList {
	switch t := i.(type) {
	default:
		panic(fmt.Sprintf("Unexpected StmtList type %T\n", t))
	case Question:
		s.Questions = append(s.Questions, i.(Question))
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

func (s StmtList) Accept(v visit.Visitor, sy interface{}) interface{} {
	return v.Visit(s, sy)
}
