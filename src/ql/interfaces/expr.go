package interfaces

type Expr interface {
	Eval(s interface{}) interface{}
	TypeCheck(TypeChecker, SymbolTable)
}
