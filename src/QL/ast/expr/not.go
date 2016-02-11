package expr

import "ql/ast/visit"

type Not struct {
	Value Expr
}

func (n Not) getValue() Expr {
	return n.Value
}

func (n Not) Eval() interface{} {
	return !n.getValue().Eval().(bool)
}

func (n Not) Accept(v visit.Visitor) interface{} {
	return v.Visit(n)
}
