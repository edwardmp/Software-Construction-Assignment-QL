package interfaces

type TypeChecker interface {
	AddEncounteredErrorForCheckType(string, error)
	GetEncountedWarnings() []error
	GetEncountedErrors() []error
	IsLabelUsed(StrLit) bool
	MarkLabelAsUsed(StrLit, VarDecl)
	VarIdForLabel(StrLit) VarId
	VarDeclIsKnown(VarDecl) bool
	MarkVarDeclAsKnown(varDecl VarDecl)
	TypeForVarDecl(varDecl VarDecl) VarType
	SetCurrentVarIdVisited(VarDecl)
	AddDependencyForCurrentlyVisitedVarDecl(VarId)
	UnsetCurrentVarIdVisited()
	GetDependencyChainForVarId(VarId) []VarId
}
