package interfaces

type SymbolTable interface {
	GetNodeForIdentifier(v VarId) interface{}
	SetNodeForIdentifier(e interface{}, v VarId)
	SaveToDisk() (interface{}, error)
}
