package expr

import (
	"ql/ast/vari"
	"ql/ast/visit"
)

type VarExpr struct {
	Identifier vari.VarId
}

// TODO implement
func (v VarExpr) Eval(symTable interface{}) interface{} {
	return false
}

func (v VarExpr) Accept(vis visit.Visitor) interface{} {
	return vis.Visit(v)
}
