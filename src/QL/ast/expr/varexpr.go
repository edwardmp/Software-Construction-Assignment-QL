package expr

import "ql/ast/vari"

type VarExpr struct {
	Identifier vari.VarId
}

func (v VarExpr) Eval(symTable interface{}) interface{} {
	return false
}
