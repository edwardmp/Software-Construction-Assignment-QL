package unaryoperatorexpr

import (
	"math"
	"ql/ast/expr"
)

type Pos struct {
	Value expr.Expr
}

func (p Pos) Eval(s interface{}) interface{} {
	return int(math.Abs(float64(p.Value.Eval(s).(int))))
}
