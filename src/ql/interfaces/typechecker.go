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
	AddDependencyForVarDecl(VarId, VarDecl)
	DependencyListForVarDeclContainsReferenceToSelf(VarDecl) bool
	GetConditionsDependentOnForVarDecl(VarDecl) []Expr
}

type TypeCheckArgs interface {
    TypeChecker() TypeChecker 
    Symbols() TypeCheckSymbols 
    CurrentVarDeclVisited() VarDecl 
    SetCurrentVarDeclVisited(VarDecl) TypeCheckArgs
    ConditionsDependentOn() []Expr 
    AddConditionDependentOn(Expr) TypeCheckArgs
}
