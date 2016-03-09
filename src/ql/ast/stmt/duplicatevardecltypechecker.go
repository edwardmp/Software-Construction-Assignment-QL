package stmt

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"ql/ast/vari"
)

type DuplicateVarDeclTypeChecker struct {
	ErrorsEncountered []error
	VarDeclEncountered map[vari.VarId]vari.VarType
}

func CheckForDuplicateVarDeclWithDiffTypes(form Form) []error {
	log.Info("Start check for duplicate var decl with different types")
	varDeclEncountered := make(map[vari.VarId]vari.VarType)
	duplicateVarDeclTypeChecker := DuplicateVarDeclTypeChecker{VarDeclEncountered: varDeclEncountered}

	form.typeCheckDuplicateDeclaration(&duplicateVarDeclTypeChecker)
	log.WithFields(log.Fields{"ErrorsEncountered": duplicateVarDeclTypeChecker.ErrorsEncountered}).Info("Ended check for duplicate var decl with different types")

	return duplicateVarDeclTypeChecker.ErrorsEncountered
}

func (f Form) typeCheckDuplicateDeclaration(duplicateVarDeclTypeChecker *DuplicateVarDeclTypeChecker) {
    f.Content.typeCheckDuplicateDeclaration(duplicateVarDeclTypeChecker)
}

func (i If) typeCheckDuplicateDeclaration(duplicateVarDeclTypeChecker *DuplicateVarDeclTypeChecker) {
    i.Body.typeCheckDuplicateDeclaration(duplicateVarDeclTypeChecker)
}

func (i IfElse) typeCheckDuplicateDeclaration(duplicateVarDeclTypeChecker *DuplicateVarDeclTypeChecker) {
    i.IfBody.typeCheckDuplicateDeclaration(duplicateVarDeclTypeChecker)
    i.ElseBody.typeCheckDuplicateDeclaration(duplicateVarDeclTypeChecker)
}

func (s StmtList) typeCheckDuplicateDeclaration(duplicateVarDeclTypeChecker *DuplicateVarDeclTypeChecker) {
    for _, question := range s.Questions {
        switch question.(type) {
            case InputQuestion:
                question.(InputQuestion).typeCheckDuplicateDeclaration(duplicateVarDeclTypeChecker)
            case ComputedQuestion:
                question.(ComputedQuestion).typeCheckDuplicateDeclaration(duplicateVarDeclTypeChecker)
        }
    }

    for _, conditional := range s.Conditionals {
        switch conditional.(type) {
        case If:
            conditional.(If).typeCheckDuplicateDeclaration(duplicateVarDeclTypeChecker)
        
        case IfElse:
            conditional.(IfElse).typeCheckDuplicateDeclaration(duplicateVarDeclTypeChecker)
        }
    }
}

func (i InputQuestion) typeCheckDuplicateDeclaration(duplicateVarDeclTypeChecker *DuplicateVarDeclTypeChecker) {
    handleQuestionForDuplicatelabelTypeCheck(i, duplicateVarDeclTypeChecker)
}

func (c ComputedQuestion) typeCheckDuplicateDeclaration(duplicateVarDeclTypeChecker *DuplicateVarDeclTypeChecker) {
    handleQuestionForDuplicatelabelTypeCheck(c, duplicateVarDeclTypeChecker)
}

func handleQuestionForDuplicatelabelTypeCheck(question Question, duplicateVarDeclTypeChecker *DuplicateVarDeclTypeChecker) {
    varDecl := question.GetVarDecl()
    labelKnown := checkIfVarDeclIsKnown(varDecl, duplicateVarDeclTypeChecker.VarDeclEncountered)

    if labelKnown && duplicateVarDeclTypeChecker.VarDeclEncountered[varDecl.Ident] != varDecl.Type {
        duplicateVarDeclTypeChecker.ErrorsEncountered = append(duplicateVarDeclTypeChecker.ErrorsEncountered, fmt.Errorf("Question redeclared with different types: %T and %T", varDecl.Type, duplicateVarDeclTypeChecker.VarDeclEncountered[varDecl.Ident]))
    } else {
        markVarDeclAsKnown(varDecl, duplicateVarDeclTypeChecker.VarDeclEncountered)
    }
}

func checkIfVarDeclIsKnown(varDecl vari.VarDecl, knownIdentifiers map[vari.VarId]vari.VarType) bool {
	if _, exists := knownIdentifiers[varDecl.Ident]; exists {
		return true
	}

	return false
}

func markVarDeclAsKnown(varDecl vari.VarDecl, knownIdentifiers map[vari.VarId]vari.VarType) {
	log.WithFields(log.Fields{"VarDecl": varDecl}).Debug("Marking VarDecl as known")
	knownIdentifiers[varDecl.Ident] = varDecl.Type
}
