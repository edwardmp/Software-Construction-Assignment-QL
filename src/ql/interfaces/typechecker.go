package interfaces

type TypeChecker interface {
	AddEncounteredError(error)
	AddEncounteredWarning(error)
	GetEncountedWarnings() []error
	GetEncountedErrors() []error
	IsLabelUsed(StrLit) bool
	MarkLabelAsUsed(StrLit, VarDecl)
	VarIdForLabel(StrLit) VarId
	VarDeclIsKnown(VarDecl) bool
	MarkVarDeclAsKnown(varDecl VarDecl)
	TypeForVarDecl(varDecl VarDecl) ValueType
	SetCurrentVarIdVisited(VarDecl)
	AddDependencyForCurrentlyVisitedVarDecl(VarId)
	UnsetCurrentVarIdVisited()
	GetDependencyChainForVarId(VarId) []VarId
	AddConditionDependentOn(Expr)
	PopLastConditionDependentOn()
	GetConditionsDependentOn() []Expr
}
