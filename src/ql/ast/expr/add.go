package expr

import (
	"ql/interfaces"
)

type Add struct {
	BinaryOperator
}

func NewAdd(lhs interfaces.Expr, rhs interfaces.Expr, sourceInfo interface{}) Add {
	return Add{NewBinaryOperator(lhs, rhs, sourceInfo)}
}

func NewAddNoSourceInfo(lhs interfaces.Expr, rhs interfaces.Expr) Add {
	return NewAdd(lhs, rhs, nil)
}
