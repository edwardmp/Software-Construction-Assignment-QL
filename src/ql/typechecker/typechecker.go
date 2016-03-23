package typechecker

import (
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

type TypeChecker struct {
	interfaces.TypeChecker
	ErrorsEncountered      []error
	WarningsEncountered    []error
	UsedLabels             map[interfaces.StrLit]interfaces.VarId
	IdentifiersEncountered map[interfaces.VarId]bool
	DependenciesPerVarId   map[interfaces.VarId][]interfaces.VarId
}

func NewTypeChecker() *TypeChecker {
	return &TypeChecker{ErrorsEncountered: make([]error, 0), WarningsEncountered: make([]error, 0), UsedLabels: make(map[interfaces.StrLit]interfaces.VarId), IdentifiersEncountered: make(map[interfaces.VarId]bool), DependenciesPerVarId: make(map[interfaces.VarId][]interfaces.VarId)}
}

func (this *TypeChecker) dependencyListForVarDecl(varDecl interfaces.VarDecl) []interfaces.VarId {
    varId := varDecl.GetIdent()

	if varId == nil {
		panic("Attempting to get dependencies of nil varId")
	}

    dependencies := make([]interfaces.VarId, 0)

    for _, directDependentVarId := range this.DependenciesPerVarId[varId] {
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

	dependencies := make([]interfaces.VarId, 0)
	for _, dependentVarId := range this.DependenciesPerVarId[varIdToObtainDependenciesFor] {
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
        if dependentVarId == varDecl.GetIdent() {
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

	for _, dependingVarIdAlreadyKnown := range this.DependenciesPerVarId[varDecl.GetIdent()] {
		// don't add it again if its already known to be a dependency
		if dependingVarIdAlreadyKnown == varIdDependentOn {
			return
		}
	}

	this.DependenciesPerVarId[varDecl.GetIdent()] = append(this.DependenciesPerVarId[varDecl.GetIdent()], varIdDependentOn)

	log.WithFields(log.Fields{"visitedVarId": varDecl.GetIdent(), "varIdDependentOn": varIdDependentOn, "resultingMap": this.DependenciesPerVarId}).Debug("Added dependency for currently visited VarId for type checking")
}

func (this *TypeChecker) AddEncounteredError(encounteredError error) {
	log.WithFields(log.Fields{"errorEncountered": encounteredError}).Info("Added encountered type checking error")

	this.ErrorsEncountered = append(this.ErrorsEncountered, encounteredError)
}

func (this *TypeChecker) AddEncounteredWarning(encounteredWarning error) {
	log.WithFields(log.Fields{"warningEncountered": encounteredWarning}).Info("Added encountered type checking warning")

	this.WarningsEncountered = append(this.WarningsEncountered, encounteredWarning)
}

func (this *TypeChecker) GetEncounteredErrors() []error {
	return this.ErrorsEncountered
}

func (this *TypeChecker) GetEncounteredWarnings() []error {
	return this.WarningsEncountered
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
	if isKnown, exists := this.IdentifiersEncountered[varDecl.GetIdent()]; exists {
		return isKnown
	}

	return false
}

func (this *TypeChecker) MarkVarIdAsKnown(varId interfaces.VarId) {
	log.WithFields(log.Fields{"VarDecl": varId}).Debug("Marking VarDecl as known")
	this.IdentifiersEncountered[varId] = true
}

// MarkVarIdAsUnknown stores that the VarId is currently unknown, if VarId remains unmarked it is a reference to a undefined question
func (this *TypeChecker) MarkVarIdAsUnknown(varId interfaces.VarId) {
	log.WithFields(log.Fields{"VarDecl": varId}).Debug("Marking VarDecl as unknown")
	this.IdentifiersEncountered[varId] = false
}

func (this *TypeChecker) GetIdentifiersEncountered() map[interfaces.VarId]bool {
	return this.IdentifiersEncountered
}
