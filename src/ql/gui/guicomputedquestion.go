package gui

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/vari"
	"ql/ast/vari/vartype"
)

type GUIComputedQuestion struct {
	GUIQuestion
	Expr  expr.Expr
	VarId vari.VarId
}

func CreateGUIComputedQuestion(label string, questionType vartype.VarType, expr expr.Expr, varId vari.VarId) GUIComputedQuestion {
	questionLabel := createLabel(label)
	questionElement := CreateDisabledInputTextField(fmt.Sprintf("%s", questionType.GetDefaultValue()))
	errorLabel := createLabel("")

	guiQuestion := GUIQuestion{questionLabel, questionElement, errorLabel}
	return GUIComputedQuestion{GUIQuestion: guiQuestion, Expr: expr, VarId: varId}
}
