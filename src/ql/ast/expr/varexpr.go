package expr

import "ql/interfaces"

type VarExpr struct {
	Identifier interfaces.VarId
	Expr
}

func NewVarExpr(identifier interfaces.VarId) VarExpr {
	return VarExpr{identifier, NewExpr()}
}

func (this VarExpr) GetIdentifier() interfaces.VarId {
	return this.Identifier
}
