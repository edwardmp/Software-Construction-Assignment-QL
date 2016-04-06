package typechecker

import (
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

type TypeChecker struct {
	interfaces.TypeChecker
	errorsEncountered    []error
	warningsEncountered  []error
	knownErrors          map[string]bool
	usedLabels           map[interfaces.StringLiteral]interfaces.VarId
	knownIdentifiers     map[interfaces.VarId]bool
	dependenciesForVarId map[interfaces.VarId][]interfaces.VarId
}

func NewTypeChecker() *TypeChecker {
	return &TypeChecker{errorsEncountered: make([]error, 0), warningsEncountered: make([]error, 0), usedLabels: make(map[interfaces.StringLiteral]interfaces.VarId), knownErrors: make(map[string]bool), knownIdentifiers: make(map[interfaces.VarId]bool), dependenciesForVarId: make(map[interfaces.VarId][]interfaces.VarId)}
}

func (this *TypeChecker) dependencyListForVarDecl(varDecl interfaces.VarDecl) []interfaces.VarId {
	varId := varDecl.VariableIdentifier()

	if varId == nil {
		panic("Attempting to get dependencies of nil varId")
	}

	var dependencies []interfaces.VarId

	for _, directDependentVarId := range this.dependenciesForVarId[varId] {
		// first add the direct dependencies as dependencies
		dependencies = append(dependencies, directDependentVarId)

		// then add as dependencies: the dependencies of the directDependency (i.e. the indirect dependencies of the VarId passed)
		dependencies = append(dependencies, this.recursivelyObtainDependenciesForVarId(directDependentVarId, varId)...)
	}

	return dependencies
}

func (this *TypeChecker) recursivelyObtainDependenciesForVarId(varIdToObtainDependenciesFor, varIdNeededToCompleteCycle interfaces.VarId) []interfaces.VarId {
	if varIdToObtainDependenciesFor == nil {
		panic("Attempting to recursively obtain dependencies for nil varIdDependentOn")
	} else if varIdNeededToCompleteCycle == nil {
		panic("Can't recursively obtain dependencies when varIdNeededToCompleteCycle is nil (no stop condition)")
	}

	// if we try to get dependencies of the currently visited VarId we've found a cycle, thus return to prevent infinite loop :)
	if varIdToObtainDependenciesFor == varIdNeededToCompleteCycle {
		return nil
	}

	var dependencies []interfaces.VarId
	for _, dependentVarId := range this.dependenciesForVarId[varIdToObtainDependenciesFor] {
		// if we encounter ourselves, we found a cyclic dependency
		if dependentVarId == varIdToObtainDependenciesFor {
			return nil
		}

		dependencies = append(dependencies, dependentVarId)
		dependencies = append(dependencies, this.recursivelyObtainDependenciesForVarId(dependentVarId, varIdNeededToCompleteCycle)...)
	}

	return dependencies
}

func (this *TypeChecker) DependencyListForVarDeclContainsReferenceToSelf(varDecl interfaces.VarDecl) bool {
	depencyListForQuestionId := this.dependencyListForVarDecl(varDecl)

	for _, dependentVarId := range depencyListForQuestionId {
		if dependentVarId == varDecl.VariableIdentifier() {
			return true
		}
	}

	return false
}

func (this *TypeChecker) AddDependencyForVarDecl(varIdDependentOn interfaces.VarId, varDecl interfaces.VarDecl) {
	// if we are not in a ComputedQuestion context, don't add any dependencies
	if varDecl == nil {
		return
	}

	if varIdDependentOn == nil {
		panic("Attempting to add VarId dependency but passed dependency is nil")
	}

	for _, dependingVarIdAlreadyKnown := range this.dependenciesForVarId[varDecl.VariableIdentifier()] {
		// don't add it again if its already known to be a dependency
		if dependingVarIdAlreadyKnown == varIdDependentOn {
			return
		}
	}

	varId := varDecl.VariableIdentifier()
	this.dependenciesForVarId[varId] = append(this.dependenciesForVarId[varId], varIdDependentOn)

	log.WithFields(log.Fields{"visitedVarId": varId, "varIdDependentOn": varIdDependentOn, "resultingMap": this.dependenciesForVarId}).Debug("Added dependency for currently visited VarId for type checking")
}

func (this *TypeChecker) AddEncounteredError(encounteredError error) {
	if encounteredError == nil {
		panic("Trying to add nil encountered error")
	}

	log.WithFields(log.Fields{"errorEncountered": encounteredError}).Info("Added encountered type checking error")

	if !this.isErrorKnown(encounteredError) {
		this.errorsEncountered = append(this.errorsEncountered, encounteredError)
		this.knownErrors[encounteredError.Error()] = true
	}
}

func (this *TypeChecker) AddEncounteredWarning(encounteredWarning error) {
	if encounteredWarning == nil {
		panic("Trying to add nil encountered warning")
	}

	log.WithFields(log.Fields{"warningEncountered": encounteredWarning}).Info("Added encountered type checking warning")

	if !this.isErrorKnown(encounteredWarning) {
		this.warningsEncountered = append(this.warningsEncountered, encounteredWarning)
		this.knownErrors[encounteredWarning.Error()] = true
	}
}

func (this *TypeChecker) isErrorKnown(errorToCheck error) bool {
	_, errorAlreadyKnown := this.knownErrors[errorToCheck.Error()]
	return errorAlreadyKnown
}

func (this *TypeChecker) EncounteredErrors() []error {
	return this.errorsEncountered
}

func (this *TypeChecker) EncounteredWarnings() []error {
	return this.warningsEncountered
}

func (this *TypeChecker) IsLabelUsed(label interfaces.StringLiteral) bool {
	_, exists := this.usedLabels[label]
	return exists
}

func (this *TypeChecker) VarIdForLabel(label interfaces.StringLiteral) interfaces.VarId {
	return this.usedLabels[label]
}

func (this *TypeChecker) MarkLabelAsUsed(label interfaces.StringLiteral, varDecl interfaces.VarDecl) {
	log.WithFields(log.Fields{"label": label}).Debug("Marking label as used")
	this.usedLabels[label] = varDecl.VariableIdentifier()
}

func (this *TypeChecker) VarDeclIsKnown(varDecl interfaces.VarDecl) bool {
	if isKnown, exists := this.knownIdentifiers[varDecl.VariableIdentifier()]; exists {
		return isKnown
	}

	return false
}

func (this *TypeChecker) MarkVarIdAsKnown(varId interfaces.VarId) {
	log.WithFields(log.Fields{"VarDecl": varId}).Debug("Marking VarDecl as known")
	this.knownIdentifiers[varId] = true
}

// MarkVarIdAsUnknown stores that the VarId is currently unknown, if VarId remains unmarked it is a reference to a undefined question
func (this *TypeChecker) MarkVarIdAsUnknown(varId interfaces.VarId) {
	log.WithFields(log.Fields{"VarDecl": varId}).Debug("Marking VarDecl as unknown")
	this.knownIdentifiers[varId] = false
}

// KnownIdentifiers returns a map where the keys are identifiers that are currently known
func (this *TypeChecker) KnownIdentifiers() map[interfaces.VarId]bool {
	return this.knownIdentifiers
}
