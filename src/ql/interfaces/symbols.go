package interfaces

type Symbols interface {
}

type TypeCheckSymbols interface {
	Symbols
	SetTypeForVarId(ValueType, VarId)
	GetTypeForVarId(VarId) ValueType
	IsTypeSetForVarId(VarId) bool
}

type VarIdValueSymbols interface {
	Symbols
	SetExprForVarId(Expr, VarId)
	GetExprForVarId(VarId) Expr
	RegisterCallback(callback func(VarIdValueSymbols))
	SaveToDisk() (interface{}, error)
}
