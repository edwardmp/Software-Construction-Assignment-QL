package unaryoperatorexpr

import "ql/ast/expr"

type Not struct {
	UnaryOperator
}

func NewNot(value expr.Expr) Not {
	return Not{UnaryOperator{Value: value}}
}

func (n Not) Eval(s interface{}) interface{} {
	return !n.Value.Eval(s).(bool)
}
