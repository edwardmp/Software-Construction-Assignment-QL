package expr

import "ql/interfaces"

type Or struct {
	BinaryOperator
}

func NewOr(lhs interfaces.Expr, rhs interfaces.Expr, sourceInfo interface{}) Or {
	return Or{NewBinaryOperator(lhs, rhs, sourceInfo)}
}

func NewOrNoSourceInfo(lhs interfaces.Expr, rhs interfaces.Expr) Or {
	return NewOr(lhs, rhs, nil)
}
