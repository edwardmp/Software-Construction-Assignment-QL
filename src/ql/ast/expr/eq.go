package expr

import "ql/interfaces"

type Eq struct {
	BinaryOperator
}

func NewEq(lhs interfaces.Expr, rhs interfaces.Expr, sourceInfo interface{}) Eq {
	return Eq{NewBinaryOperator(lhs, rhs, sourceInfo)}
}

func NewEqNoSourceInfo(lhs interfaces.Expr, rhs interfaces.Expr) Eq {
	return NewEq(lhs, rhs, nil)
}
