package typechecker

import (
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

type TypeChecker struct {
	interfaces.TypeChecker
	ErrorsEncounteredForCheckType map[string][]error
	UsedLabels                    map[interfaces.StrLit]interfaces.VarId
	KnownIdentifiers              map[interfaces.VarId]interfaces.VarType
}

func NewTypeChecker() TypeChecker {
	return TypeChecker{ErrorsEncounteredForCheckType: make(map[string][]error), UsedLabels: make(map[interfaces.StrLit]interfaces.VarId), KnownIdentifiers: make(map[interfaces.VarId]interfaces.VarType)}
}

func (t *TypeChecker) AddEncounteredErrorForCheckType(checkType string, encounteredError error) {
	log.WithFields(log.Fields{"checkType": checkType, "errorEncountered": encounteredError}).Info("Added encountered error for check type")

	t.ErrorsEncounteredForCheckType[checkType] = append(t.ErrorsEncounteredForCheckType[checkType], encounteredError)
}

func (t *TypeChecker) GetEncountedErrorsForCheckType(checkType string) []error {
	return t.ErrorsEncounteredForCheckType[checkType]
}

func (t *TypeChecker) IsLabelUsed(label interfaces.StrLit) bool {
	if _, exists := t.UsedLabels[label]; exists {
		return true
	}

	return false
}

func (t *TypeChecker) VarIdForLabel(label interfaces.StrLit) interfaces.VarId {
	return t.UsedLabels[label]
}

func (t *TypeChecker) MarkLabelAsUsed(label interfaces.StrLit, varDecl interfaces.VarDecl) {
	log.WithFields(log.Fields{"label": label}).Debug("Marking label as used")
	t.UsedLabels[label] = varDecl.GetIdent()
}

func (t *TypeChecker) VarDeclIsKnown(varDecl interfaces.VarDecl) bool {
	if _, exists := t.KnownIdentifiers[varDecl.GetIdent()]; exists {
		return true
	}

	return false
}

func (t *TypeChecker) MarkVarDeclAsKnown(varDecl interfaces.VarDecl) {
	log.WithFields(log.Fields{"VarDecl": varDecl}).Debug("Marking VarDecl as known")
	t.KnownIdentifiers[varDecl.GetIdent()] = varDecl.GetType()
}

func (t *TypeChecker) TypeForVarDecl(varDecl interfaces.VarDecl) interfaces.VarType {
	return t.KnownIdentifiers[varDecl.GetIdent()]
}
