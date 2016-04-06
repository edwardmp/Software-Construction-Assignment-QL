package gui

import (
	"ql/interfaces"
)

type GUIComputedQuestion struct {
	*GUIQuestion
	Expr  interfaces.Expr
	VarID interfaces.VarID
}

// newGUIComputedQuestion is a constructr method returning a new GUIComputedQuestion
func newGUIComputedQuestion(label string, questionType interfaces.ValueType, expr interfaces.Expr, varID interfaces.VarID) *GUIComputedQuestion {
	guiQuestion := createDisabledGUIQuestion(label, questionType, nil)
	return &GUIComputedQuestion{GUIQuestion: guiQuestion, Expr: expr, VarID: varID}
}
