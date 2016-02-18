package expr

import (
	"ql/ast/vari"
	"ql/ast/visit"
)

type VarExpr struct {
	Identifier vari.VarId
	Expr       Expr
}

func (v VarExpr) GetIdentifier() vari.VarId {
	return v.Identifier
}

func (v VarExpr) GetExpr() Expr {
	return v.Expr
}

func (v VarExpr) SetExpr(e Expr) Expr {
	v.Expr = e
	return v
}

func (v VarExpr) Eval() interface{} {
	return v.Expr
}

func (v VarExpr) Accept(vis visit.Visitor, s interface{}) interface{} {
	return vis.Visit(v, s)
}
