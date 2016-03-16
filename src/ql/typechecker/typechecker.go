package typechecker

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

type TypeChecker struct {
	interfaces.TypeChecker
	ErrorsEncounteredForCheckType map[string][]error
	UsedLabels                    map[interfaces.StrLit]interfaces.VarId
	KnownIdentifiers              map[interfaces.VarId]interfaces.VarType
	CurrentVarIdVisited           interfaces.VarId
	DependenciesPerVarId          map[interfaces.VarId][]interfaces.VarId
}

func NewTypeChecker() TypeChecker {
	return TypeChecker{ErrorsEncounteredForCheckType: make(map[string][]error), UsedLabels: make(map[interfaces.StrLit]interfaces.VarId), KnownIdentifiers: make(map[interfaces.VarId]interfaces.VarType), DependenciesPerVarId: make(map[interfaces.VarId][]interfaces.VarId)}
}

func (this *TypeChecker) SetCurrentVarIdVisited(varDeclVisiting interfaces.VarDecl) {
	varIdVisiting := varDeclVisiting.GetIdent()

	if varIdVisiting == nil {
		panic("Attempting to set current VarId visited to nil")
	}

	log.WithFields(log.Fields{"visitedVarId": varIdVisiting}).Debug("Set currently visited VarId for typechecking")

	this.CurrentVarIdVisited = varIdVisiting
	this.DependenciesPerVarId[this.CurrentVarIdVisited] = make([]interfaces.VarId, 0)
}

func (this *TypeChecker) UnsetCurrentVarIdVisited() {
	this.CurrentVarIdVisited = nil
}

func (this *TypeChecker) GetDependencyChainForVarId(varId interfaces.VarId) []interfaces.VarId {
	if varId == nil {
		panic("Attempting to get dependencies of nil varId")
	}

	dependencies := make([]interfaces.VarId, 0)
	dependencies = append(dependencies, this.CurrentVarIdVisited)

	for _, directDependency := range this.DependenciesPerVarId[this.CurrentVarIdVisited] {
		dependencies = append(dependencies, directDependency)

		// add as dependencies: the dependencies of the directDependency (i.e. the indirect dependencies of CurrentVarIdVisited) in a transitive fashion
		dependencies = append(dependencies, this.getDependenciesOfDirectDependenciesForVarId(directDependency)...)
	}

	fmt.Println(dependencies)

	return dependencies
}

func (this *TypeChecker) getDependenciesOfDirectDependenciesForVarId(varId interfaces.VarId) []interfaces.VarId {
	// if we try to get dependencies of the currently visited var id we've found a cycle
	if varId == this.CurrentVarIdVisited {
		return []interfaces.VarId{}
	}

	dependencies := make([]interfaces.VarId, 0)
	for _, directDependency := range this.DependenciesPerVarId[varId] {
		dependencies = append(dependencies, directDependency)
		dependencies = append(dependencies, this.getDependenciesOfDirectDependenciesForVarId(directDependency)...)
	}

	return dependencies
}

func (this *TypeChecker) AddDependencyForCurrentlyVisitedVarDecl(dependingVarId interfaces.VarId) {
	// if we are not in a ComputedQuestion context, don't add any dependencies
	if this.CurrentVarIdVisited == nil {
		return
	}

	if dependingVarId == nil {
		panic("Attempting to add VarId dependency but passed dependency is nil")
	}

	for _, dependingVarIdAlreadyKnown := range this.DependenciesPerVarId[this.CurrentVarIdVisited] {
		// don't add it again if its already known to be a dependency
		if dependingVarIdAlreadyKnown == dependingVarId {
			return
		}
	}

	this.DependenciesPerVarId[this.CurrentVarIdVisited] = append(this.DependenciesPerVarId[this.CurrentVarIdVisited], dependingVarId)

	log.WithFields(log.Fields{"visitedVarId": this.CurrentVarIdVisited, "dependingVarId": dependingVarId, "resultingMap": this.DependenciesPerVarId}).Debug("Added dependency for currently visited VarId for type checking")
}

func (this *TypeChecker) AddEncounteredErrorForCheckType(checkType string, encounteredError error) {
	log.WithFields(log.Fields{"checkType": checkType, "errorEncountered": encounteredError}).Info("Added encountered error for check type")

	this.ErrorsEncounteredForCheckType[checkType] = append(this.ErrorsEncounteredForCheckType[checkType], encounteredError)
}

func (this *TypeChecker) GetEncountedErrors() []error {
	errorsCollected := make([]error, 0)
	errorsCollected = append(errorsCollected, this.ErrorsEncounteredForCheckType["InvalidOperandsDifferentTypes"]...)
	errorsCollected = append(errorsCollected, this.ErrorsEncounteredForCheckType["ReferenceToUndefinedQuestion"]...)
	errorsCollected = append(errorsCollected, this.ErrorsEncounteredForCheckType["InvalidOperationOnOperands"]...)
	errorsCollected = append(errorsCollected, this.ErrorsEncounteredForCheckType["NonBoolConditionals"]...)
	errorsCollected = append(errorsCollected, this.ErrorsEncounteredForCheckType["CyclicalDependencies"]...)

	return errorsCollected
}

func (this *TypeChecker) GetEncountedWarnings() []error {
	return this.ErrorsEncounteredForCheckType["DuplicateLabels"]
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
