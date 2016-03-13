package vari

import (
	"ql/interfaces"
)

func (this VarDecl) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {
	this.Ident.TypeCheck(typeChecker, symbols)
	this.Type.TypeCheck(typeChecker, symbols)
}

func (this VarId) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {
}

func (this IntType) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {
}

func (this StringType) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {
}

func (this BoolType) TypeCheck(typeChecker interfaces.TypeChecker, symbols interfaces.Symbols) {
}
