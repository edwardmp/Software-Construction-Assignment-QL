package interfaces

type TypeCheckSymbols interface {
	SetTypeForVarId(ValueType, VarId)
	GetTypeForVarId(VarId) ValueType
	IsTypeSetForVarId(VarId) bool
}

type VarIdValueSymbols interface {
	SetExprForVarId(Expr, VarId)
	GetExprForVarId(VarId) Expr
	RegisterCallback(callback func())
	SaveToDisk() (interface{}, error)
}
