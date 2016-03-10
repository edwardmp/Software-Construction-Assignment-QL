package binaryoperatorexpr

import "ql/ast/expr"

type Div struct {
	BinaryOperator
}

func NewDiv(lhs expr.Expr, rhs expr.Expr) Div {
	return Div{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (d Div) Eval(s interface{}) interface{} {
	return d.Lhs.Eval(s).(int) / d.Rhs.Eval(s).(int)
}
