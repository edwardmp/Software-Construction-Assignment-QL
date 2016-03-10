package expr

import "ql/interfaces"

type GT struct {
	BinaryOperator
}

func NewGT(lhs interfaces.Expr, rhs interfaces.Expr) GT {
	return GT{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (g GT) Eval(s interface{}) interface{} {
	switch g.Lhs.Eval(s).(type) {
	case int:
		return g.Lhs.Eval(s).(int) > g.Rhs.Eval(s).(int)
	case string:
		return g.Lhs.Eval(s).(string) > g.Rhs.Eval(s).(string)
	}

	return nil
}
