package expr

import (
	"ql/ast/vari"
	"ql/ast/visit"
	"ql/env"
)

type VarExpr struct {
	Identifier vari.VarId
}

func (v VarExpr) GetIdentifier() vari.VarId {
	return v.Identifier
}

func (v VarExpr) Eval(s interface{}) interface{} {
	symbolTable := s.(env.SymbolTable)
	return symbolTable.GetNodeForIdentifier(v.Identifier).(Expr).Eval(s)
}

func (v VarExpr) Accept(vis visit.Visitor, s interface{}) interface{} {
	return vis.Visit(v, s)
}
