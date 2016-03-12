package expr

import "ql/interfaces"

type Neg struct {
	UnaryOperator
}

func NewNeg(value interfaces.Expr, sourceInfo interface{}) Neg {
	return Neg{NewUnaryOperator(value, sourceInfo)}
}

func NewNegNoSourceInfo(value interfaces.Expr) Neg {
	return NewNeg(value, nil)
}
