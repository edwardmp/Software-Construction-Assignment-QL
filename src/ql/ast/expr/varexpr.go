package expr

import (
	"fmt"
	"ql/interfaces"
	"ql/symboltable"
)

type VarExpr struct {
	Identifier interfaces.VarId
}

func (v VarExpr) Eval(s interface{}) interface{} {
	symbolTable, castOK := s.(symboltable.SymbolTable)

	if !castOK {
		fmt.Print(s)
		panic("No symbol table passed to Eval Varinterfaces.Expr")
	}

	if node := symbolTable.GetNodeForIdentifier(v.Identifier); node != nil {
		return node.(interfaces.Expr).Eval(s)
	}

	return nil
}
