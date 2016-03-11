package expr

import "ql/interfaces"

type VarExpr struct {
	Identifier interfaces.VarId
}

func (v VarExpr) GetIdentifier() interfaces.VarId {
	return v.Identifier
}
