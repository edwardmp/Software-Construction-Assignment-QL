package interfaces

type ASTNode interface {
	Node
	Accept(Visitor, Symbols) interface{}
	TypeCheck(TypeChecker, Symbols)
}
