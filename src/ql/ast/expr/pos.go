package expr

import "ql/interfaces"

type Pos struct {
	UnaryOperator
}

func NewPos(value interfaces.Expr, sourceInfo interface{}) Pos {
	return Pos{NewUnaryOperator(value, sourceInfo)}
}

func NewPosNoSourceInfo(value interfaces.Expr) Pos {
	return NewPos(value, nil)
}
