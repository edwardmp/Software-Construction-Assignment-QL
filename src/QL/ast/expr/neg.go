package expr

import (
	"math"
	"ql/ast/visit"
)

type Neg struct {
	Value Expr
}

func (n Neg) getValue() Expr {
	return n.Value
}

func (n Neg) Eval() interface{} {
	return int(math.Abs(float64(n.getValue().Eval().(int))) * -1)
}

func (n Neg) Accept(v visit.Visitor) interface{} {
	return v.Visit(n)
}
