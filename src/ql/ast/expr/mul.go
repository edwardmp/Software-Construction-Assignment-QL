package expr

import "ql/interfaces"

type Mul struct {
	BinaryOperator
}

func NewMul(lhs interfaces.Expr, rhs interfaces.Expr) Mul {
	return Mul{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}
