package expr

import "ql/interfaces"

type Div struct {
	BinaryOperator
}

func NewDiv(lhs interfaces.Expr, rhs interfaces.Expr, sourceInfo interface{}) Div {
	return Div{NewBinaryOperator(lhs, rhs, sourceInfo)}
}

func NewDivNoSourceInfo(lhs interfaces.Expr, rhs interfaces.Expr) Div {
	return NewDiv(lhs, rhs, nil)
}
