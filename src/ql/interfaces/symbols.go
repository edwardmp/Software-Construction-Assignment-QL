package interfaces

type Symbols interface {
	GetNodeForIdentifier(v VarId) interface{}
	SetNodeForIdentifier(e interface{}, v VarId)
	SaveToDisk() (interface{}, error)
	RegisterCallback(callback func(Symbols))
}
