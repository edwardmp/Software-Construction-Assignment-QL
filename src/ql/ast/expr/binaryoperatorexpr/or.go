package binaryoperatorexpr

import "ql/ast/expr"

type Or struct {
	BinaryOperator
}

func NewOr(lhs expr.Expr, rhs expr.Expr) Or {
	return Or{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (o Or) Eval(s interface{}) interface{} {
	return o.Lhs.Eval(s).(bool) || o.Rhs.Eval(s).(bool)
}
