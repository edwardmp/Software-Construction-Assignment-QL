package stmt

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"ql/ast/expr/litexpr"
	"ql/ast/vari"
)

type DuplicateLabelTypeChecker struct {
	ErrorsEncountered []error
	LabelsEncountered map[litexpr.StrLit]vari.VarId
}

func CheckForDuplicateLabels(form Form) []error {
	log.Info("Start check for duplicate labels")
	labelsEncountered := make(map[litexpr.StrLit]vari.VarId)
	duplicateLabelTypeChecker := DuplicateLabelTypeChecker{LabelsEncountered: labelsEncountered}

	form.typeCheckDuplicateLabels(&duplicateLabelTypeChecker)
	log.WithFields(log.Fields{"ErrorsEncountered": duplicateLabelTypeChecker.ErrorsEncountered}).Info("Ended check for duplicate labels")

	return duplicateLabelTypeChecker.ErrorsEncountered
}

func (f Form) typeCheckDuplicateLabels(duplicateLabelTypeChecker *DuplicateLabelTypeChecker) {
	f.Content.typeCheckDuplicateLabels(duplicateLabelTypeChecker)
}

func (i If) typeCheckDuplicateLabels(duplicateLabelTypeChecker *DuplicateLabelTypeChecker) {
    i.Body.typeCheckDuplicateLabels(duplicateLabelTypeChecker)
}

func (i IfElse) typeCheckDuplicateLabels(duplicateLabelTypeChecker *DuplicateLabelTypeChecker) {
    i.IfBody.typeCheckDuplicateLabels(duplicateLabelTypeChecker)
    i.ElseBody.typeCheckDuplicateLabels(duplicateLabelTypeChecker)
}

func (s StmtList) typeCheckDuplicateLabels(duplicateLabelTypeChecker *DuplicateLabelTypeChecker) {
    for _, question := range s.Questions {
        switch question.(type) {
            case InputQuestion:
                question.(InputQuestion).typeCheckDuplicateLabels(duplicateLabelTypeChecker)
            case ComputedQuestion:
                question.(ComputedQuestion).typeCheckDuplicateLabels(duplicateLabelTypeChecker)
        }
    }

    for _, conditional := range s.Conditionals {
        switch conditional.(type) {
        case If:
            conditional.(If).typeCheckDuplicateLabels(duplicateLabelTypeChecker)
        
        case IfElse:
            conditional.(IfElse).typeCheckDuplicateLabels(duplicateLabelTypeChecker)
        }
    }
}

func (i InputQuestion) typeCheckDuplicateLabels(duplicateLabelTypeChecker *DuplicateLabelTypeChecker) {
	handleQuestion(i, duplicateLabelTypeChecker)
}

func (c ComputedQuestion) typeCheckDuplicateLabels(duplicateLabelTypeChecker *DuplicateLabelTypeChecker) {
    handleQuestion(c, duplicateLabelTypeChecker)
}

func handleQuestion(question Question, duplicateLabelTypeChecker *DuplicateLabelTypeChecker) {
    labelKnown := checkIfLabelIsUsed(question.GetLabel(), duplicateLabelTypeChecker.LabelsEncountered)

    if labelKnown {
        duplicateLabelTypeChecker.ErrorsEncountered = append(duplicateLabelTypeChecker.ErrorsEncountered, fmt.Errorf("Label \"%s\" already used for question with identifier %s, using again for question with identifier %s", question.GetLabel(), duplicateLabelTypeChecker.LabelsEncountered[question.GetLabel()], question.GetVarDecl().Ident))
    } else {
        markLabelAsUsed(question.GetLabel(), question.GetVarDecl(), duplicateLabelTypeChecker.LabelsEncountered)
    }
}

func checkIfLabelIsUsed(label litexpr.StrLit, usedLabels map[litexpr.StrLit]vari.VarId) bool {
	if _, exists := usedLabels[label]; exists {
		return true
	}

	return false
}

func markLabelAsUsed(label litexpr.StrLit, varDecl vari.VarDecl, usedLabels map[litexpr.StrLit]vari.VarId) {
	log.WithFields(log.Fields{"label": label}).Debug("Marking label as used")
	usedLabels[label] = varDecl.Ident
}
