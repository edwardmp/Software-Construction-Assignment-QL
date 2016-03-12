package expr

import "ql/interfaces"

type NEq struct {
	BinaryOperator
}

func NewNEq(lhs interfaces.Expr, rhs interfaces.Expr, sourceInfo interface{}) NEq {
	return NEq{NewBinaryOperator(lhs, rhs, sourceInfo)}
}

func NewNEqNoSourceInfo(lhs interfaces.Expr, rhs interfaces.Expr) NEq {
	return NewNEq(lhs, rhs, nil)
}
