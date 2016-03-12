package interfaces

import "ql/token"

type Node interface {
	Accept(Visitor, interface{}) interface{}
	TypeCheck(TypeChecker, SymbolTable)
	GetSourceInfo() token.Pos
}
