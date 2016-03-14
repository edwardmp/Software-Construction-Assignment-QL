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

func (this *TypeChecker) AddEncounteredErrorForCheckType(checkType string, encounteredError error) {
	log.WithFields(log.Fields{"checkType": checkType, "errorEncountered": encounteredError}).Info("Added encountered error for check type")

	this.ErrorsEncounteredForCheckType[checkType] = append(this.ErrorsEncounteredForCheckType[checkType], encounteredError)
}

func (this *TypeChecker) GetEncountedErrorsForCheckType(checkType string) []error {
	return this.ErrorsEncounteredForCheckType[checkType]
}

func (this *TypeChecker) IsLabelUsed(label interfaces.StrLit) bool {
	if _, exists := this.UsedLabels[label]; exists {
		return true
	}

	return false
}

func (this *TypeChecker) VarIdForLabel(label interfaces.StrLit) interfaces.VarId {
	return this.UsedLabels[label]
}

func (this *TypeChecker) MarkLabelAsUsed(label interfaces.StrLit, varDecl interfaces.VarDecl) {
	log.WithFields(log.Fields{"label": label}).Debug("Marking label as used")
	this.UsedLabels[label] = varDecl.GetIdent()
}

func (this *TypeChecker) VarDeclIsKnown(varDecl interfaces.VarDecl) bool {
	if _, exists := this.KnownIdentifiers[varDecl.GetIdent()]; exists {
		return true
	}

	return false
}

func (this *TypeChecker) MarkVarDeclAsKnown(varDecl interfaces.VarDecl) {
	log.WithFields(log.Fields{"VarDecl": varDecl}).Debug("Marking VarDecl as known")
	this.KnownIdentifiers[varDecl.GetIdent()] = varDecl.GetType()
}

func (this *TypeChecker) TypeForVarDecl(varDecl interfaces.VarDecl) interfaces.VarType {
	return this.KnownIdentifiers[varDecl.GetIdent()]
}
