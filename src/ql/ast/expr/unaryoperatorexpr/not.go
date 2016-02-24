package unaryoperatorexpr

import (
	"ql/ast/expr"
	"ql/ast/visit"
)

type Not struct {
	Value expr.Expr
}

func (n Not) GetValue() expr.Expr {
	return n.Value
}

func (n Not) Eval(s interface{}) interface{} {
	return !n.GetValue().Eval(s).(bool)
}

func (n Not) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(n, s)
}
