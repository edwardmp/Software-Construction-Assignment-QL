package expr

import (
	"fmt"
	"ql/ast/vari"
	"ql/ast/visit"
	"ql/symboltable"
)

type VarExpr struct {
	Identifier vari.VarId
}

func (v VarExpr) GetIdentifier() vari.VarId {
	return v.Identifier
}

func (v VarExpr) Eval(s interface{}) interface{} {
	symbolTable, castOK := s.(symboltable.SymbolTable)

	if !castOK {
		fmt.Print(s)
		panic("No symbol table passed to Eval VarExpr")
	}

	return symbolTable.GetNodeForIdentifier(v.Identifier).(Expr).Eval(s)
}

func (v VarExpr) Accept(vis visit.Visitor, s interface{}) interface{} {
	return vis.Visit(v, s)
}
