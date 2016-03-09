package unaryoperatorexpr

import (
	"ql/ast/expr"
)

type Not struct {
	Value expr.Expr
}

func (n Not) Eval(s interface{}) interface{} {
	return !n.Value.Eval(s).(bool)
}
