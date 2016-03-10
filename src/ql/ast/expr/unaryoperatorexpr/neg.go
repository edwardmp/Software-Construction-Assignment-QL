package unaryoperatorexpr

import (
	"math"
	"ql/ast/expr"
)

type Neg struct {
	UnaryOperator
}

func NewNeg(value expr.Expr) Neg {
	return Neg{UnaryOperator{Value: value}}
}

func (n Neg) Eval(s interface{}) interface{} {
	return int(math.Abs(float64(n.Value.Eval(s).(int))) * -1)
}
