package expr

import "math"

type Neg struct {
	Value Expr
}

func (neg Neg) Eval() interface{} {
	return int(math.Abs(float64(neg.Value.Eval().(int))) * -1)
}
