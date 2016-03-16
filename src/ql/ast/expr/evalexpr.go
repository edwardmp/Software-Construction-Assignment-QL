package expr

import (
	"math"
	"ql/interfaces"
)

func (this Add) Eval(s interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(s).(int) + this.Rhs.Eval(s).(int)
}

func (this And) Eval(s interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(s).(bool) && this.Rhs.Eval(s).(bool)
}

func (this Div) Eval(s interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(s).(int) / this.Rhs.Eval(s).(int)
}

func (this Eq) Eval(s interfaces.VarIdValueSymbols) interface{} {
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

func (this GEq) Eval(s interfaces.VarIdValueSymbols) interface{} {
	switch this.Lhs.Eval(s).(type) {
	case int:
		return this.Lhs.Eval(s).(int) >= this.Rhs.Eval(s).(int)
	case string:
		return this.Lhs.Eval(s).(string) >= this.Rhs.Eval(s).(string)
	}

	return nil
}

func (this GT) Eval(s interfaces.VarIdValueSymbols) interface{} {
	switch this.Lhs.Eval(s).(type) {
	case int:
		return this.Lhs.Eval(s).(int) > this.Rhs.Eval(s).(int)
	case string:
		return this.Lhs.Eval(s).(string) > this.Rhs.Eval(s).(string)
	}

	return nil
}

func (this LEq) Eval(s interfaces.VarIdValueSymbols) interface{} {
	switch this.Lhs.Eval(s).(type) {
	case int:
		return this.Lhs.Eval(s).(int) <= this.Rhs.Eval(s).(int)
	case string:
		return this.Lhs.Eval(s).(string) <= this.Rhs.Eval(s).(string)
	}

	return nil
}

func (this LT) Eval(s interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(s).(int) < this.Rhs.Eval(s).(int)
}

func (this Mul) Eval(s interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(s).(int) * this.Rhs.Eval(s).(int)
}

func (this Neg) Eval(s interfaces.VarIdValueSymbols) interface{} {
	return int(math.Abs(float64(this.Value.Eval(s).(int))) * -1)
}

func (this NEq) Eval(s interfaces.VarIdValueSymbols) interface{} {
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

func (this Not) Eval(s interfaces.VarIdValueSymbols) interface{} {
	return !this.Value.Eval(s).(bool)
}

func (this Or) Eval(s interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(s).(bool) || this.Rhs.Eval(s).(bool)
}

func (this Pos) Eval(s interfaces.VarIdValueSymbols) interface{} {
	return int(math.Abs(float64(this.Value.Eval(s).(int))))
}

func (this IntLit) Eval(s interfaces.VarIdValueSymbols) interface{} {
	return this.Value
}

func (this StrLit) Eval(s interfaces.VarIdValueSymbols) interface{} {
	return this.Value
}

func (this BoolLit) Eval(s interfaces.VarIdValueSymbols) interface{} {
	return this.Value
}

func (this Sub) Eval(s interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(s).(int) - this.Rhs.Eval(s).(int)
}

func (this VarExpr) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	if symbols == nil {
		panic("No symbol table passed to Eval VarExpr")
	}

	if referencedExpr := symbols.GetExprForVarId(this.Identifier); referencedExpr != nil {
		return referencedExpr.Eval(symbols)
	}

	return nil
}

func (this Expr) Eval(s interfaces.VarIdValueSymbols) interface{} {
	panic("Expr struct Eval method not overridden")

	return nil
}
