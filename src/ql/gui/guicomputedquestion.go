package gui

import (
	"ql/ast/expr"
)

type GUIComputedQuestion struct {
	GUIQuestion
	Expr expr.Expr
}

func CreateGUIComputedQuestion(label string, elementValueString string, expr expr.Expr) GUIComputedQuestion {
	questionLabel := createQuestionLabel(label)
	questionElement := CreateDisabledInputTextField(elementValueString)
	errorLabel := createQuestionLabel("")

	guiQuestion := GUIQuestion{questionLabel, questionElement, errorLabel}
	return GUIComputedQuestion{GUIQuestion: guiQuestion, Expr: expr}
}
