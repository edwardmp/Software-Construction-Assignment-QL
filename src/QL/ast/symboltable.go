package ast

import (
	"ql/ast/expr"
	"ql/ast/vari"
)

type SymbolTable struct {
	Table map[vari.VarId]expr.Expr
}

func (s SymbolTable) getExprForIdentifier(v vari.VarId) expr.Expr {
	return s.Table[v]
}
