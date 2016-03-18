package expr

import (
	"math"
	"ql/interfaces"
)

func (this Add) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(symbols).(int) + this.Rhs.Eval(symbols).(int)
}

func (this And) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(symbols).(bool) && this.Rhs.Eval(symbols).(bool)
}

func (this Div) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	denominator := this.Rhs.Eval(symbols).(int)

	// since default value of int question will be set to zero, handle it here to prevent division by zero panic
	if denominator == 0 {
		return 0
	}

	return this.Lhs.Eval(symbols).(int) / denominator
}

func (this Eq) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	var returnValue bool

	// we assume both sides are of equal types as this should have been caught during type checking
	switch lhsEvalValue := this.Lhs.Eval(symbols).(type) {
	case int:
		returnValue = lhsEvalValue == this.Rhs.Eval(symbols).(int)
	case bool:
		returnValue = lhsEvalValue == this.Rhs.Eval(symbols).(bool)
	case string:
		returnValue = lhsEvalValue == this.Rhs.Eval(symbols).(string)
	default:
		panic("Eval result is of unknown type")
	}

	return returnValue
}

func (this GEq) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(symbols).(int) >= this.Rhs.Eval(symbols).(int)
}

func (this GT) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(symbols).(int) > this.Rhs.Eval(symbols).(int)
}

func (this LEq) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(symbols).(int) <= this.Rhs.Eval(symbols).(int)
}

func (this LT) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(symbols).(int) < this.Rhs.Eval(symbols).(int)
}

func (this Mul) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(symbols).(int) * this.Rhs.Eval(symbols).(int)
}

func (this Neg) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return int(math.Abs(float64(this.Value.Eval(symbols).(int))) * -1)
}

func (this NEq) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	var returnValue bool

	// we assume both sides are of equal types as this should have been caught during type checking
	switch lhsEvalValue := this.Lhs.Eval(symbols).(type) {
	case int:
		returnValue = lhsEvalValue != this.Rhs.Eval(symbols).(int)
	case bool:
		returnValue = lhsEvalValue != this.Rhs.Eval(symbols).(bool)
	case string:
		returnValue = lhsEvalValue != this.Rhs.Eval(symbols).(string)
	default:
		panic("Eval result is of unknown type")
	}

	return returnValue
}

func (this Not) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return !this.Value.Eval(symbols).(bool)
}

func (this Or) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(symbols).(bool) || this.Rhs.Eval(symbols).(bool)
}

func (this Pos) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return int(math.Abs(float64(this.Value.Eval(symbols).(int))))
}

func (this IntLit) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return this.Value
}

func (this StrLit) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return this.Value
}

func (this BoolLit) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return this.Value
}

func (this Sub) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	return this.Lhs.Eval(symbols).(int) - this.Rhs.Eval(symbols).(int)
}

func (this VarExpr) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	if symbols == nil {
		panic("No symbols passed to Eval VarExpr")
	}

	referencedExpr := symbols.GetExprForVarId(this.Identifier)
	if referencedExpr == nil {
		panic("VarExpr refers to nil expression")
	}

	return referencedExpr.Eval(symbols)
}

func (this Expr) Eval(symbols interfaces.VarIdValueSymbols) interface{} {
	panic("Expr Eval method not overridden")

	return nil
}
