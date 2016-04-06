package gui

import (
	"ql/interfaces"
)

type GUIInputQuestion struct {
	*GUIQuestion
}

// newGUIInputQuestion is a construct method return a new GUIInputQuestion of the supplied question type
func newGUIInputQuestion(label string, questionType interfaces.ValueType, callback func(interfaces.Expr, error)) *GUIInputQuestion {
	return &GUIInputQuestion{GUIQuestion: createEnabledGUIQuestion(label, questionType, callback)}
}
