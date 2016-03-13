package expr

import (
	"math"
	"ql/interfaces"
)

func (this Add) Eval(s interfaces.Symbols) interface{} {
	return this.Lhs.Eval(s).(int) + this.Rhs.Eval(s).(int)
}

func (this And) Eval(s interfaces.Symbols) interface{} {
	return this.Lhs.Eval(s).(bool) && this.Rhs.Eval(s).(bool)
}

func (this BoolLit) Eval(s interfaces.Symbols) interface{} {
	return bool(this.Value)
}

func (this Div) Eval(s interfaces.Symbols) interface{} {
	return this.Lhs.Eval(s).(int) / this.Rhs.Eval(s).(int)
}

func (this Eq) Eval(s interfaces.Symbols) interface{} {
	switch this.Lhs.Eval(s).(type) {
	case int:
		return this.Lhs.Eval(s).(int) == this.Rhs.Eval(s).(int)
	case bool:
		return this.Lhs.Eval(s).(bool) == this.Rhs.Eval(s).(bool)
	case string:
		return this.Lhs.Eval(s).(string) == this.Rhs.Eval(s).(string)
	}

	return nil
}

func (this GEq) Eval(s interfaces.Symbols) interface{} {
	switch this.Lhs.Eval(s).(type) {
	case int:
		return this.Lhs.Eval(s).(int) >= this.Rhs.Eval(s).(int)
	case string:
		return this.Lhs.Eval(s).(string) >= this.Rhs.Eval(s).(string)
	}

	return nil
}

func (this GT) Eval(s interfaces.Symbols) interface{} {
	switch this.Lhs.Eval(s).(type) {
	case int:
		return this.Lhs.Eval(s).(int) > this.Rhs.Eval(s).(int)
	case string:
		return this.Lhs.Eval(s).(string) > this.Rhs.Eval(s).(string)
	}

	return nil
}

func (this IntLit) Eval(s interfaces.Symbols) interface{} {
	return this.Value
}

func (this LEq) Eval(s interfaces.Symbols) interface{} {
	switch this.Lhs.Eval(s).(type) {
	case int:
		return this.Lhs.Eval(s).(int) <= this.Rhs.Eval(s).(int)
	case string:
		return this.Lhs.Eval(s).(string) <= this.Rhs.Eval(s).(string)
	}

	return nil
}

func (this LT) Eval(s interfaces.Symbols) interface{} {
	return this.Lhs.Eval(s).(int) < this.Rhs.Eval(s).(int)
}

func (this Mul) Eval(s interfaces.Symbols) interface{} {
	return this.Lhs.Eval(s).(int) * this.Rhs.Eval(s).(int)
}

func (this Neg) Eval(s interfaces.Symbols) interface{} {
	return int(math.Abs(float64(this.Value.Eval(s).(int))) * -1)
}

func (this NEq) Eval(s interfaces.Symbols) interface{} {
	switch this.Lhs.Eval(s).(type) {
	case int:
		return this.Lhs.Eval(s).(int) != this.Rhs.Eval(s).(int)
	case bool:
		return this.Lhs.Eval(s).(bool) != this.Rhs.Eval(s).(bool)
	case string:
		return this.Lhs.Eval(s).(string) != this.Rhs.Eval(s).(string)
	}

	return nil
}

func (this Not) Eval(s interfaces.Symbols) interface{} {
	return !this.Value.Eval(s).(bool)
}

func (this Or) Eval(s interfaces.Symbols) interface{} {
	return this.Lhs.Eval(s).(bool) || this.Rhs.Eval(s).(bool)
}

func (this Pos) Eval(s interfaces.Symbols) interface{} {
	return int(math.Abs(float64(this.Value.Eval(s).(int))))
}

func (this StrLit) Eval(s interfaces.Symbols) interface{} {
	return string(this.Value)
}

func (this Sub) Eval(s interfaces.Symbols) interface{} {
	return this.Lhs.Eval(s).(int) - this.Rhs.Eval(s).(int)
}

func (this VarExpr) Eval(symbols interfaces.Symbols) interface{} {
	if symbols == nil {
		panic("No symbol table passed to Eval VarExpr")
	}

	if referencedExpr := symbols.GetNodeForIdentifier(this.Identifier); referencedExpr != nil {
		return referencedExpr.(interfaces.Expr).Eval(symbols)
	}

	return nil
}

func (this Expr) Eval(s interfaces.Symbols) interface{} {
	panic("Expr struct Eval method not overridden")
	return nil
}
