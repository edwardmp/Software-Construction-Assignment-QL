package expr

import "ql/interfaces"

type Not struct {
	UnaryOperator
}

func NewNot(value interfaces.Expr, sourceInfo interface{}) Not {
	return Not{NewUnaryOperator(value, sourceInfo)}
}

func NewNotNoSourceInfo(value interfaces.Expr) Not {
	return NewNot(value, nil)
}
