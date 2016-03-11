package gui

import (
	"ql/interfaces"
)

type GUIInputQuestion struct {
	GUIQuestion
}

func CreateGUIInputQuestion(label string, questionType interfaces.VarType, callback func(interface{}, error)) GUIInputQuestion {
	return GUIInputQuestion{GUIQuestion: CreateEnabledGUIQuestion(label, questionType, callback)}
}
