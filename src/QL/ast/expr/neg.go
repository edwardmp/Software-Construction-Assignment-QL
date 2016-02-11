package expr

import "math"

type Neg struct {
	Value Expr
}

func (neg Neg) getValue() Expr {
	return neg.Value
}

func (neg Neg) Eval() interface{} {
	return int(math.Abs(float64(neg.getValue().Eval().(int))) * -1)
}
