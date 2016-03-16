package gui

import (
	"ql/interfaces"
)

type GUIComputedQuestion struct {
	*GUIQuestion
	Expr  interfaces.Expr
	VarId interfaces.VarId
}

// CreateGUIComputedQuestion is a constructr method returning a new GUIComputedQuestion
func CreateGUIComputedQuestion(label string, questionType interfaces.ValueType, expr interfaces.Expr, varId interfaces.VarId) *GUIComputedQuestion {
	guiQuestion := CreateDisabledGUIQuestion(label, questionType, nil)
	return &GUIComputedQuestion{GUIQuestion: guiQuestion, Expr: expr, VarId: varId}
}
