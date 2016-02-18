package unaryoperatorexpr

import (
	"math"
	"ql/ast/expr"
	"ql/ast/visit"
)

type Neg struct {
	Value expr.Expr
}

func (n Neg) GetValue() expr.Expr {
	return n.Value
}

func (n Neg) Eval() interface{} {
	return int(math.Abs(float64(n.GetValue().Eval().(int))) * -1)
}

func (n Neg) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(n, s)
}
