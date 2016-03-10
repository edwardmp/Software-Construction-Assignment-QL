package binaryoperatorexpr

import "ql/ast/expr"

type Mul struct {
	BinaryOperator
}

func NewMul(lhs expr.Expr, rhs expr.Expr) Mul {
	return Mul{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (m Mul) Eval(s interface{}) interface{} {
	return m.Lhs.Eval(s).(int) * m.Rhs.Eval(s).(int)
}
