package expr

import "ql/interfaces"

type Mul struct {
	BinaryOperator
}

func NewMul(lhs interfaces.Expr, rhs interfaces.Expr, sourceInfo interface{}) Mul {
	return Mul{NewBinaryOperator(lhs, rhs, sourceInfo)}
}

func NewMulNoSourceInfo(lhs interfaces.Expr, rhs interfaces.Expr) Mul {
	return NewMul(lhs, rhs, nil)
}
