package expr

import "ql/interfaces"

type LEq struct {
	BinaryOperator
}

func NewLEq(lhs interfaces.Expr, rhs interfaces.Expr) LEq {
	return LEq{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (l LEq) Eval(s interface{}) interface{} {
	switch l.Lhs.Eval(s).(type) {
	case int:
		return l.Lhs.Eval(s).(int) <= l.Rhs.Eval(s).(int)
	case string:
		return l.Lhs.Eval(s).(string) <= l.Rhs.Eval(s).(string)
	}

	return nil
}
