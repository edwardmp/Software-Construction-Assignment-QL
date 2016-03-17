package interfaces

type TypeChecker interface {
	AddEncounteredError(error)
	AddEncounteredWarning(error)
	GetEncountedWarnings() []error
	GetEncountedErrors() []error
	IsLabelUsed(StrLit) bool
	MarkLabelAsUsed(StrLit, VarDecl)
	VarIdForLabel(StrLit) VarId
	MarkVarIdAsKnown(VarId)
	MarkVarIdAsUnknown(VarId)
	GetIdentifiersEncountered() map[VarId]bool
	SetCurrentVarIdVisited(VarDecl)
	AddDependencyForCurrentlyVisitedVarDecl(VarId)
	UnsetCurrentVarIdVisited()
	GetDependencyChainForVarId(VarId) []VarId
	AddConditionDependentOn(Expr)
	PopLastConditionDependentOn()
	GetConditionsDependentOn() []Expr
}
