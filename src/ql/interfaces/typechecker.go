package interfaces

type TypeChecker interface {
	AddEncounteredError(error)
	AddEncounteredWarning(error)
	EncounteredWarnings() []error
	EncounteredErrors() []error
	IsLabelUsed(StrLit) bool
	MarkLabelAsUsed(StrLit, VarDecl)
	VarIdForLabel(StrLit) VarId
	MarkVarIdAsKnown(VarId)
	MarkVarIdAsUnknown(VarId)
	IdentifiersEncountered() map[VarId]bool
	AddDependencyForVarDecl(VarId, VarDecl)
	DependencyListForVarDeclContainsReferenceToSelf(VarDecl) bool
	ConditionsDependentOnForVarDecl(VarDecl) []Expr
}

type TypeCheckArgs interface {
	TypeChecker() TypeChecker
	Symbols() TypeCheckSymbols
	CurrentVarDeclVisited() VarDecl
	SetCurrentVarDeclVisited(VarDecl) TypeCheckArgs
	ConditionsDependentOn() []Expr
	AddConditionDependentOn(Expr) TypeCheckArgs
}
