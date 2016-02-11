package stmt

import "fmt"

type StmtList struct {
	Questions    []Question
	Conditionals []Conditional
}

func (s StmtList) AddToCorrectSlice(i interface{}) StmtList {
	switch t := i.(type) {
	default:
		panic(fmt.Sprintf("unexpected StmtList type %T\n", t))
	case Question:
		s.Questions = append(s.Questions, i.(Question))
	case If:
		s.Conditionals = append(s.Conditionals, i.(If))
	case IfElse:
		s.Conditionals = append(s.Conditionals, i.(IfElse))
	}

	return s
}
