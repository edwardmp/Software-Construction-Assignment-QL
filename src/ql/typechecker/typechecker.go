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
	CurrentVarIdVisited    interfaces.VarId
	DependenciesPerVarId   map[interfaces.VarId][]interfaces.VarId
	ConditionsDependentOn  []interfaces.Expr
}

func NewTypeChecker() *TypeChecker {
	return &TypeChecker{ErrorsEncountered: make([]error, 0), WarningsEncountered: make([]error, 0), UsedLabels: make(map[interfaces.StrLit]interfaces.VarId), IdentifiersEncountered: make(map[interfaces.VarId]bool), DependenciesPerVarId: make(map[interfaces.VarId][]interfaces.VarId)}
}

func (this *TypeChecker) AddConditionDependentOn(condition interfaces.Expr) {
	if condition == nil {
		panic("Attempting to add nil dependent condition")
	}

	log.WithFields(log.Fields{"conditionDependentOn": condition}).Debug("Add condition dependent on")

	this.ConditionsDependentOn = append(this.ConditionsDependentOn, condition)
}

func (this *TypeChecker) GetConditionsDependentOn() []interfaces.Expr {
	conditionsDependentOn := this.ConditionsDependentOn

	log.WithFields(log.Fields{"conditionsDependentOn": conditionsDependentOn}).Debug("Getting conditions dependent on")

	return conditionsDependentOn
}

func (this *TypeChecker) PopLastConditionDependentOn() {
	conditionsDependentOn := this.ConditionsDependentOn

	if len(conditionsDependentOn) == 0 {
		panic("Attempting to pop from conditions dependent on but not aware of any")
	}

	log.WithFields(log.Fields{"conditionDependentOnPopped": this.ConditionsDependentOn[len(this.ConditionsDependentOn)-1]}).Debug("Pop last condition dependent on")

	// remove last element
	this.ConditionsDependentOn = conditionsDependentOn[:len(conditionsDependentOn)-1]
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

	return dependencies
}

func (this *TypeChecker) getDependenciesOfDirectDependenciesForVarId(varId interfaces.VarId) []interfaces.VarId {
	if varId == nil {
		panic("Attempting to get dependencies of nil varId")
	}

	// if we try to get dependencies of the currently visited VarId we've found a cycle
	if varId == this.CurrentVarIdVisited {
		return nil
	}

	dependencies := make([]interfaces.VarId, 0)
	for _, directDependency := range this.DependenciesPerVarId[varId] {
		// if we encounter ourselves, we found a cyclic dependency
		if directDependency == varId {
			return nil
		}

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

// MarkVarIdAsUnknown stores that the VarId is currently unknown
// If the VarId is not marked as known later it is a reference to a undefined question
func (this *TypeChecker) MarkVarIdAsUnknown(varId interfaces.VarId) {
	log.WithFields(log.Fields{"VarDecl": varId}).Debug("Marking VarDecl as unknown")
	this.IdentifiersEncountered[varId] = false
}

func (this *TypeChecker) GetIdentifiersEncountered() map[interfaces.VarId]bool {
	return this.IdentifiersEncountered
}
