package expr

import (
	"math"
	"ql/ast/visit"
)

type Pos struct {
	Value Expr
}

func (p Pos) getValue() Expr {
	return p.Value
}

func (p Pos) Eval() interface{} {
	return int(math.Abs(float64(p.getValue().Eval().(int))))
}

func (p Pos) Accept(v visit.Visitor) interface{} {
	return v.Visit(p)
}
