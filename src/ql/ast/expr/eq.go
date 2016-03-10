package expr

import "ql/interfaces"

type Eq struct {
	BinaryOperator
}

func NewEq(lhs interfaces.Expr, rhs interfaces.Expr) Eq {
	return Eq{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}
