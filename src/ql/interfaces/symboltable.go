package interfaces

import "ql/ast/vari"

type SymbolTable interface {
    GetNodeForIdentifier(v vari.VarId) interface{} 
    SetNodeForIdentifier(e interface{}, v vari.VarId)
    SaveToDisk() (interface{}, error) 
}