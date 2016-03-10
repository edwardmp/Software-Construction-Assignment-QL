package binaryoperatorexpr

import "ql/ast/expr"

type NEq struct {
	BinaryOperator
}

func NewNEq(lhs expr.Expr, rhs expr.Expr) NEq {
	return NEq{BinaryOperator{Lhs: lhs, Rhs: rhs}}
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
