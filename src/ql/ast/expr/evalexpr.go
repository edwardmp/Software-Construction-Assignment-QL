package expr

import (
	"fmt"
	"math"
	"ql/interfaces"
)

func (a Add) Eval(s interface{}) interface{} {
	return a.Lhs.Eval(s).(int) + a.Rhs.Eval(s).(int)
}

func (a And) Eval(s interface{}) interface{} {
	return a.Lhs.Eval(s).(bool) && a.Rhs.Eval(s).(bool)
}

func (b BoolLit) Eval(s interface{}) interface{} {
	return bool(b.Value)
}

func (d Div) Eval(s interface{}) interface{} {
	return d.Lhs.Eval(s).(int) / d.Rhs.Eval(s).(int)
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

func (g GEq) Eval(s interface{}) interface{} {
	switch g.Lhs.Eval(s).(type) {
	case int:
		return g.Lhs.Eval(s).(int) >= g.Rhs.Eval(s).(int)
	case string:
		return g.Lhs.Eval(s).(string) >= g.Rhs.Eval(s).(string)
	}

	return nil
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

func (i IntLit) Eval(s interface{}) interface{} {
	return i.Value
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

func (l LT) Eval(s interface{}) interface{} {
	return l.Lhs.Eval(s).(int) < l.Rhs.Eval(s).(int)
}

func (m Mul) Eval(s interface{}) interface{} {
	return m.Lhs.Eval(s).(int) * m.Rhs.Eval(s).(int)
}

func (n Neg) Eval(s interface{}) interface{} {
	return int(math.Abs(float64(n.Value.Eval(s).(int))) * -1)
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

func (n Not) Eval(s interface{}) interface{} {
	return !n.Value.Eval(s).(bool)
}

func (o Or) Eval(s interface{}) interface{} {
	return o.Lhs.Eval(s).(bool) || o.Rhs.Eval(s).(bool)
}

func (p Pos) Eval(s interface{}) interface{} {
	return int(math.Abs(float64(p.Value.Eval(s).(int))))
}

func (s StrLit) Eval(sy interface{}) interface{} {
	return string(s.Value)
}

func (s Sub) Eval(sy interface{}) interface{} {
	return s.Lhs.Eval(sy).(int) - s.Rhs.Eval(sy).(int)
}

func (v VarExpr) Eval(s interface{}) interface{} {
	symbolTable, castOK := s.(interfaces.SymbolTable)

	if !castOK {
		fmt.Print(s)
		panic("No symbol table passed to Eval Varinterfaces.Expr")
	}

	if node := symbolTable.GetNodeForIdentifier(v.Identifier); node != nil {
		return node.(interfaces.Expr).Eval(s)
	}

	return nil
}
