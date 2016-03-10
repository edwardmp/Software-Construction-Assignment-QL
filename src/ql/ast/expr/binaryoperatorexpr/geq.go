package binaryoperatorexpr

import "ql/ast/expr"

type GEq struct {
	BinaryOperator
}

func NewGEq(lhs expr.Expr, rhs expr.Expr) GEq {
	return GEq{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (g GEq) Eval(s interface{}) interface{} {
	return g.Lhs.Eval(s).(int) >= g.Rhs.Eval(s).(int)
}
