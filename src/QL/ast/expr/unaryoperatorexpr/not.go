package unaryoperatorexpr

import ("ql/ast/visit"
"ql/ast/expr")

type Not struct {
	Value expr.Expr
}

func (n Not) GetValue() expr.Expr {
	return n.Value
}

func (n Not) Eval() interface{} {
	return !n.GetValue().Eval().(bool)
}

func (n Not) Accept(v visit.Visitor) interface{} {
	return v.Visit(n)
}
