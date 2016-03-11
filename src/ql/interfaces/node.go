package interfaces

type Node interface {
	Accept(Visitor, interface{}) interface{}
	TypeCheck(TypeChecker, SymbolTable)
}
