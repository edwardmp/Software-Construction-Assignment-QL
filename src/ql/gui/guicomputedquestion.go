package gui

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/vari"
)

type GUIComputedQuestion struct {
	GUIQuestion
	Expr  expr.Expr
	VarId vari.VarId
}

func CreateGUIComputedQuestion(label string, questionType vari.VarType, expr expr.Expr, varId vari.VarId) GUIComputedQuestion {
	guiQuestion := GUIQuestion{label, questionType, nil}
	return GUIComputedQuestion{GUIQuestion: guiQuestion, Expr: expr, VarId: varId}
}
