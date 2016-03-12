package expr

import "ql/interfaces"

type VarExpr struct {
	Identifier interfaces.VarId
	Expr
}

func NewVarExpr(identifier interfaces.VarId, sourceInfo interface{}) VarExpr {
	return VarExpr{identifier, NewExpr(sourceInfo)}
}

func NewVarExprNoSourceInfo(identifier interfaces.VarId) VarExpr {
	return NewVarExpr(identifier, nil)
}

func (v VarExpr) GetIdentifier() interfaces.VarId {
	return v.Identifier
}
