package expr

import "ql/interfaces"

type Eq struct {
	BinaryOperator
}

func NewEq(lhs interfaces.Expr, rhs interfaces.Expr) Eq {
	return Eq{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (e Eq) Eval(s interface{}) interface{} {
	switch e.Lhs.Eval(s).(type) {
	case int:
		return e.Lhs.Eval(s).(int) == e.Rhs.Eval(s).(int)
	case bool:
		return e.Lhs.Eval(s).(bool) == e.Rhs.Eval(s).(bool)
	case string:
		return e.Lhs.Eval(s).(string) == e.Rhs.Eval(s).(string)
	}

	return nil
}
