package expr

import "ql/interfaces"

type LEq struct {
	BinaryOperator
}

func NewLEq(lhs interfaces.Expr, rhs interfaces.Expr, sourceInfo interface{}) LEq {
	return LEq{NewBinaryOperator(lhs, rhs, sourceInfo)}
}

func NewLEqNoSourceInfo(lhs interfaces.Expr, rhs interfaces.Expr) LEq {
	return NewLEq(lhs, rhs, nil)
}
