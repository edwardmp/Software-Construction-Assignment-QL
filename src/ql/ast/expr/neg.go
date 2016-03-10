package expr

import (
	"math"
	"ql/interfaces"
)

type Neg struct {
	UnaryOperator
}

func NewNeg(value interfaces.Expr) Neg {
	return Neg{UnaryOperator{Value: value}}
}

func (n Neg) Eval(s interface{}) interface{} {
	return int(math.Abs(float64(n.Value.Eval(s).(int))) * -1)
}
