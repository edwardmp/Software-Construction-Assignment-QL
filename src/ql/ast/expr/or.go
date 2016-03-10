package expr

import "ql/interfaces"

type Or struct {
	BinaryOperator
}

func NewOr(lhs interfaces.Expr, rhs interfaces.Expr) Or {
	return Or{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}
