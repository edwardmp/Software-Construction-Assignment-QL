package expr

import "ql/interfaces"

type GEq struct {
	BinaryOperator
}

func NewGEq(lhs interfaces.Expr, rhs interfaces.Expr, sourceInfo interface{}) GEq {
	return GEq{NewBinaryOperator(lhs, rhs, sourceInfo)}
}

func NewGEqNoSourceInfo(lhs interfaces.Expr, rhs interfaces.Expr) GEq {
	return NewGEq(lhs, rhs, nil)
}
