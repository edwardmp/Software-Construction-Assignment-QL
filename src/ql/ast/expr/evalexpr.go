package expr

import "ql/interfaces"

func (this Add) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Lhs().Eval(symbols).Add(this.Rhs().Eval(symbols))
}

func (this And) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Lhs().Eval(symbols).And(this.Rhs().Eval(symbols))
}

func (this Div) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Lhs().Eval(symbols).Div(this.Rhs().Eval(symbols))
}

func (this Eq) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Lhs().Eval(symbols).Eq(this.Rhs().Eval(symbols))
}

func (this GEq) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Lhs().Eval(symbols).GEq(this.Rhs().Eval(symbols))
}

func (this GT) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Lhs().Eval(symbols).GT(this.Rhs().Eval(symbols))
}

func (this LEq) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Lhs().Eval(symbols).LEq(this.Rhs().Eval(symbols))
}

func (this LT) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Lhs().Eval(symbols).LT(this.Rhs().Eval(symbols))
}

func (this Mul) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Lhs().Eval(symbols).Mul(this.Rhs().Eval(symbols))
}

func (this Neg) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Value().Eval(symbols).Neg()
}

func (this NEq) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Lhs().Eval(symbols).NEq(this.Rhs().Eval(symbols))
}

func (this Not) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Value().Eval(symbols).Not()
}

func (this Or) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Lhs().Eval(symbols).Or(this.Rhs().Eval(symbols))
}

func (this Pos) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Value().Eval(symbols).Pos()
}

func (this IntegerLiteral) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Value()
}

func (this StringLiteral) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Value()
}

func (this BoolLiteral) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Value()
}

func (this Sub) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Lhs().Eval(symbols).Sub(this.Rhs().Eval(symbols))
}

func (this VarExpr) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	if symbols == nil {
		panic("No symbols passed to Eval VarExpr")
	}

	referencedExpr := symbols.ExprForVarID(this.Identifier())
	if referencedExpr == nil {
		panic("VarExpr refers to nil expression")
	}

	return referencedExpr.Eval(symbols)
}

func (this Expr) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	panic("Expr Eval method not overridden")

	return nil
}
