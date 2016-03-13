package gui

import (
	"ql/interfaces"
)

type GUIInputQuestion struct {
	*GUIQuestion
}

// CreateGUIInputQuestion is a construct method return a new GUIInputQuestion of the supplied question type.
func CreateGUIInputQuestion(label string, questionType interfaces.VarType, callback func(interface{}, error)) *GUIInputQuestion {
	return &GUIInputQuestion{GUIQuestion: CreateEnabledGUIQuestion(label, questionType, callback)}
}
