package expr

import "ql/interfaces"

type NEq struct {
	BinaryOperator
}

func NewNEq(lhs interfaces.Expr, rhs interfaces.Expr) NEq {
	return NEq{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (n NEq) Eval(s interface{}) interface{} {
	switch n.Lhs.Eval(s).(type) {
	case int:
		return n.Lhs.Eval(s).(int) != n.Rhs.Eval(s).(int)
	case bool:
		return n.Lhs.Eval(s).(bool) != n.Rhs.Eval(s).(bool)
	case string:
		return n.Lhs.Eval(s).(string) != n.Rhs.Eval(s).(string)
	}

	return nil
}
