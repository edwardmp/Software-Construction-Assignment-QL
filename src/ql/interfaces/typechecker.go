package interfaces

type TypeChecker interface {
	AddEncounteredErrorForCheckType(string, error)
	GetEncountedErrorsForCheckType(string) []error
	IsLabelUsed(StrLit) bool
	MarkLabelAsUsed(StrLit, VarDecl)
	VarIdForLabel(StrLit) VarId
	VarDeclIsKnown(VarDecl) bool
	MarkVarDeclAsKnown(varDecl VarDecl)
	TypeForVarDecl(varDecl VarDecl) VarType
}
