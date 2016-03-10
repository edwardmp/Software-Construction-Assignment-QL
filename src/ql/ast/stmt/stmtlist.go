package stmt

import (
	"fmt"
	"ql/interfaces"
)

type StmtList struct {
	Questions    []interfaces.Question
	Conditionals []interfaces.Conditional
}

func (s StmtList) AddToCorrectSlice(i interface{}) StmtList {
	switch t := i.(type) {
	default:
		panic(fmt.Sprintf("Unexpected StmtList type %T\n", t))
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
