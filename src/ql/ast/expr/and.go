package expr

import "ql/interfaces"

type And struct {
	BinaryOperator
}

func NewAnd(lhs interfaces.Expr, rhs interfaces.Expr, sourceInfo interface{}) And {
	return And{NewBinaryOperator(lhs, rhs, sourceInfo)}
}

func NewAndNoSourceInfo(lhs interfaces.Expr, rhs interfaces.Expr) And {
	return NewAnd(lhs, rhs, nil)
}
