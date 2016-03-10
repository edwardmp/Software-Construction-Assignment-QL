package expr

import "ql/interfaces"

type Not struct {
	UnaryOperator
}

func NewNot(value interfaces.Expr) Not {
	return Not{UnaryOperator{Value: value}}
}

func (n Not) Eval(s interface{}) interface{} {
	return !n.Value.Eval(s).(bool)
}
