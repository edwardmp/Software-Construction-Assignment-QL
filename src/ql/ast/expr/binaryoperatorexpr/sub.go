package binaryoperatorexpr

import "ql/ast/expr"

type Sub struct {
	BinaryOperator
}

func NewSub(lhs expr.Expr, rhs expr.Expr) Sub {
	return Sub{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (s Sub) Eval(sy interface{}) interface{} {
	return s.Lhs.Eval(sy).(int) - s.Rhs.Eval(sy).(int)
}
