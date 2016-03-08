package typechecker

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"ql/ast/expr/lit"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/ast/visit"
)

type DuplicateLabelTypeChecker struct {
	LabelsEncountered map[lit.StrLit]vari.VarId
	visit.Visitor
	WarningsEncountered []error
}

func CheckForDuplicateLabels(form stmt.Form) []error {
	log.Info("Start check for duplicate labels")
	labelsEncountered := make(map[lit.StrLit]vari.VarId)
	duplicateLabelChecker := DuplicateLabelTypeChecker{LabelsEncountered: labelsEncountered}

	duplicateLabelChecker.Visit(form, nil)
	log.WithFields(log.Fields{"WarningsEncountered": duplicateLabelChecker.WarningsEncountered}).Info("Ended check for duplicate labels")

	return duplicateLabelChecker.WarningsEncountered
}

func (v *DuplicateLabelTypeChecker) Visit(t interface{}, s interface{}) interface{} {
	switch t.(type) {
	default:
		log.WithFields(log.Fields{"Node": fmt.Sprintf("%T", t)}).Debug("Ignoring unhandled node type")
	case stmt.Form:
		log.Debug("Visit Form")
		t.(stmt.Form).Content.Accept(v, s)
	case stmt.StmtList:
		log.Debug("Visit StmtList")

		for _, question := range t.(stmt.StmtList).Questions {
			question.Accept(v, s)
		}

		for _, conditional := range t.(stmt.StmtList).Conditionals {
			conditional.Accept(v, s)
		}
	case stmt.InputQuestion, stmt.ComputedQuestion:
		log.Debug("Visit Question")

		question := t.(stmt.Question)
		labelKnown := checkIfLabelIsUsed(question.GetLabel(), v.LabelsEncountered)

		if labelKnown {
			v.WarningsEncountered = append(v.WarningsEncountered, fmt.Errorf("Label \"%s\" already used for question with identifier %s, using again for question with identifier %s", question.GetLabel(), v.LabelsEncountered[question.GetLabel()], question.GetVarDecl().Ident))
		} else {
			markLabelAsUsed(question.GetLabel(), question.GetVarDecl(), v.LabelsEncountered)
		}
	case stmt.If:
		log.Debug("Visit If")
		t.(stmt.If).Body.Accept(v, s)
	case stmt.IfElse:
		log.Debug("Visit IfElse")
		t.(stmt.IfElse).IfBody.Accept(v, s)
		t.(stmt.IfElse).ElseBody.Accept(v, s)
	}

	return v
}

func checkIfLabelIsUsed(label lit.StrLit, usedLabels map[lit.StrLit]vari.VarId) bool {
	if _, exists := usedLabels[label]; exists {
		return true
	}

	return false
}

func markLabelAsUsed(label lit.StrLit, varDecl vari.VarDecl, usedLabels map[lit.StrLit]vari.VarId) {
	log.WithFields(log.Fields{"label": label}).Debug("Marking label as used")
	usedLabels[label] = varDecl.Ident
}
