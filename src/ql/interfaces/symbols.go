package interfaces

type TypeCheckSymbols interface {
	SetTypeForVarId(ValueType, VarId)
	TypeForVarId(VarId) ValueType
	IsTypeSetForVarId(VarId) bool
}

type VarIdValueSymbols interface {
	SetExprForVarId(Expr, VarId)
	ExprForVarId(VarId) Expr
	RegisterCallback(callback func())
	SaveToDisk() (interface{}, error)
}
