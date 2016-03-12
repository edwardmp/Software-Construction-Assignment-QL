package expr

import "ql/interfaces"

type Sub struct {
	BinaryOperator
}

func NewSub(lhs interfaces.Expr, rhs interfaces.Expr, sourceInfo interface{}) Sub {
	return Sub{NewBinaryOperator(lhs, rhs, sourceInfo)}
}

func NewSubNoSourceInfo(lhs interfaces.Expr, rhs interfaces.Expr) Sub {
	return NewSub(lhs, rhs, nil)
}
