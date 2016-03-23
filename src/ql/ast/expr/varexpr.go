package expr

import "ql/interfaces"

type VarExpr struct {
	identifier interfaces.VarId
	Expr
}

func NewVarExpr(identifier interfaces.VarId) VarExpr {
	return VarExpr{identifier: identifier, Expr: NewExpr()}
}

func (this VarExpr) Identifier() interfaces.VarId {
	return this.identifier
}
