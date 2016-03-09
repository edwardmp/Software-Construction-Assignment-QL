package binaryoperatorexpr

import (
	"ql/ast/expr"
)

type NEq struct {
	Lhs, Rhs expr.Expr
}

func (n NEq) Eval(s interface{}) interface{} {
	switch n.Lhs.Eval(s).(type) {
	case int:
		return n.Lhs.Eval(s).(int) != n.Rhs.Eval(s).(int)
	case bool:
		return n.Lhs.Eval(s).(bool) != n.Rhs.Eval(s).(bool)
	default:
		panic("NEq error: comparing unknown types")
	}
}
