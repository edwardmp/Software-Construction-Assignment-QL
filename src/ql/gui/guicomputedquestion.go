package gui

import (
	"ql/interfaces"
)

type GUIComputedQuestion struct {
	*GUIQuestion
	Expr  interfaces.Expr
	VarId interfaces.VarId
}

// createGUIComputedQuestion is a constructr method returning a new GUIComputedQuestion
func createGUIComputedQuestion(label string, questionType interfaces.ValueType, expr interfaces.Expr, varId interfaces.VarId) *GUIComputedQuestion {
	guiQuestion := createDisabledGUIQuestion(label, questionType, nil)
	return &GUIComputedQuestion{GUIQuestion: guiQuestion, Expr: expr, VarId: varId}
}
