package expr

import "ql/interfaces"

type LT struct {
	BinaryOperator
}

func NewLT(lhs interfaces.Expr, rhs interfaces.Expr, sourceInfo interface{}) LT {
	return LT{NewBinaryOperator(lhs, rhs, sourceInfo)}
}

func NewLTNoSourceInfo(lhs interfaces.Expr, rhs interfaces.Expr) LT {
	return NewLT(lhs, rhs, nil)
}
