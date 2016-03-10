package expr

import "ql/interfaces"

type Or struct {
	BinaryOperator
}

func NewOr(lhs interfaces.Expr, rhs interfaces.Expr) Or {
	return Or{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (o Or) Eval(s interface{}) interface{} {
	return o.Lhs.Eval(s).(bool) || o.Rhs.Eval(s).(bool)
}
