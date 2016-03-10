package expr

import "ql/interfaces"

type GEq struct {
	BinaryOperator
}

func NewGEq(lhs interfaces.Expr, rhs interfaces.Expr) GEq {
	return GEq{BinaryOperator{Lhs: lhs, Rhs: rhs}}
}

func (g GEq) Eval(s interface{}) interface{} {
	switch g.Lhs.Eval(s).(type) {
	case int:
		return g.Lhs.Eval(s).(int) >= g.Rhs.Eval(s).(int)
	case string:
		return g.Lhs.Eval(s).(string) >= g.Rhs.Eval(s).(string)
	}

	return nil
}
