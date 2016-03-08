package typechecker

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/ast/vari/vartype"
	"ql/ast/visit"
)

type DuplicateVarDeclTypeChecker struct {
	VarDeclEncountered map[vari.VarId]vartype.VarType
	visit.Visitor
	ErrorsEncountered []error
}

func CheckForDuplicateVarDeclWithDiffTypes(form stmt.Form) {
	log.Info("Start check for duplicate var decl with different types")
	varDeclEncountered := make(map[vari.VarId]vartype.VarType)
	DuplicateVarDeclTypeChecker := DuplicateVarDeclTypeChecker{VarDeclEncountered: varDeclEncountered}

	DuplicateVarDeclTypeChecker.Visit(form, nil)
	log.WithFields(log.Fields{"ErrorsEncountered": DuplicateVarDeclTypeChecker.ErrorsEncountered}).Info("Ended check for duplicate var decl with diff types")
}

func (v *DuplicateVarDeclTypeChecker) Visit(t interface{}, s interface{}) interface{} {
	switch t.(type) {
	default:
		log.WithFields(log.Fields{"Node": fmt.Sprintf("%T", t)}).Debug("Ignoring unhandled node type")
	case stmt.Form:
		log.Debug("Visit Form")
		t.(stmt.Form).Identifier.Accept(v, s)
		t.(stmt.Form).Content.Accept(v, s)
	case vari.VarDecl:
		log.Debug("Visit VarDecl")
		varDecl := t.(vari.VarDecl)
		ident := varDecl.Ident
		labelKnown := checkIfVarDeclIsKnown(varDecl, v.VarDeclEncountered)

		if labelKnown && v.VarDeclEncountered[ident] != varDecl.Type {
			v.ErrorsEncountered = append(v.ErrorsEncountered, fmt.Errorf("Question redeclared with different types: %T and %T", varDecl.Type, v.VarDeclEncountered[ident]))
		} else {
			markVarDeclAsKnown(varDecl, v.VarDeclEncountered)
		}

		t.(vari.VarDecl).Ident.Accept(v, s)
	case stmt.StmtList:
		log.Debug("Visit StmtList")

		for _, question := range t.(stmt.StmtList).Questions {
			question.Accept(v, s)
		}

		for _, conditional := range t.(stmt.StmtList).Conditionals {
			conditional.Accept(v, s)
		}
	case stmt.InputQuestion:
		log.Debug("Visit InputQuestion")
		t.(stmt.InputQuestion).VarDecl.Accept(v, s)
	case stmt.ComputedQuestion:
		log.Debug("Visit ComputedQuestion")
		t.(stmt.ComputedQuestion).VarDecl.Accept(v, s)
	case stmt.If:
		log.Debug("Visit If")
		t.(stmt.If).Body.Accept(v, s)
	case stmt.IfElse:
		log.Debug("Visit IfElse")
		t.(stmt.IfElse).IfBody.Accept(v, s)
		t.(stmt.IfElse).ElseBody.Accept(v, s)
	}

	return nil
}

func checkIfVarDeclIsKnown(varDecl vari.VarDecl, knownIdentifiers map[vari.VarId]vartype.VarType) bool {
	if _, exists := knownIdentifiers[varDecl.Ident]; exists {
		return true
	}

	return false
}

func markVarDeclAsKnown(varDecl vari.VarDecl, knownIdentifiers map[vari.VarId]vartype.VarType) {
	log.WithFields(log.Fields{"VarDecl": varDecl}).Debug("Marking VarDecl as known")
	knownIdentifiers[varDecl.Ident] = varDecl.Type
}
