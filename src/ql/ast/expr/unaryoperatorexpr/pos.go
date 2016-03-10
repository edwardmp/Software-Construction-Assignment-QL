package unaryoperatorexpr

import (
	"math"
	"ql/ast/expr"
)

type Pos struct {
	UnaryOperator
}

func NewPos(value expr.Expr) Pos {
	return Pos{UnaryOperator{Value: value}}
}

func (p Pos) Eval(s interface{}) interface{} {
	return int(math.Abs(float64(p.Value.Eval(s).(int))))
}
