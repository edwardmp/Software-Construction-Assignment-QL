package unaryoperatorexpr

import (
	"math"
	"ql/ast/expr"
)

type Neg struct {
	Value expr.Expr
}

func (n Neg) Eval(s interface{}) interface{} {
	return int(math.Abs(float64(n.Value.Eval(s).(int))) * -1)
}