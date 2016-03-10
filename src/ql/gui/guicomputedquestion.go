package gui

import (
	"ql/interfaces"
)

type GUIComputedQuestion struct {
	GUIQuestion
	Expr  interfaces.Expr
	VarId interfaces.VarId
}

func CreateGUIComputedQuestion(label string, questionType interfaces.VarType, expr interfaces.Expr, varId interfaces.VarId) GUIComputedQuestion {
	guiQuestion := CreateGUIQuestion(label, questionType, nil)
	return GUIComputedQuestion{GUIQuestion: guiQuestion, Expr: expr, VarId: varId}
}
