package expr

import "ql/interfaces"

type Add struct {
	BinaryOperator
}

func NewAdd(lhs interfaces.Expr, rhs interfaces.Expr) Add {
	return Add{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}
