package expr

import "ql/interfaces"

type GT struct {
	BinaryOperator
}

func NewGT(lhs interfaces.Expr, rhs interfaces.Expr, sourceInfo interface{}) GT {
	return GT{NewBinaryOperator(lhs, rhs, sourceInfo)}
}

func NewGTNoSourceInfo(lhs interfaces.Expr, rhs interfaces.Expr) GT {
	return NewGT(lhs, rhs, nil)
}
