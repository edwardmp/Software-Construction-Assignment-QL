package unaryoperatorexpr

import (
	"math"
	"ql/ast/visit"
    "ql/ast/expr"
)

type Pos struct {
	Value expr.Expr
}

func (p Pos) GetValue() expr.Expr {
	return p.Value
}

func (p Pos) Eval() interface{} {
	return int(math.Abs(float64(p.GetValue().Eval().(int))))
}

func (p Pos) Accept(v visit.Visitor) interface{} {
	return v.Visit(p)
}
