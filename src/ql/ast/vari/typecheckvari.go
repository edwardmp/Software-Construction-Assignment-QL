package vari

import (
	"ql/interfaces"
)

func (v VarDecl) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
	v.Ident.TypeCheck(typeChecker, symbolTable)
	v.Type.TypeCheck(typeChecker, symbolTable)
}

func (v VarId) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
}

func (i IntType) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
}

func (s StringType) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
}

func (b BoolType) TypeCheck(typeChecker interfaces.TypeChecker, symbolTable interfaces.SymbolTable) {
}
