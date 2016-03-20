package expr

import (
	"ql/interfaces"
)

type BinaryOperator struct {
	Lhs, Rhs interfaces.Expr
	Expr
}

func NewBinaryOperator(lhs interfaces.Expr, rhs interfaces.Expr) BinaryOperator {
	return BinaryOperator{Lhs: lhs, Rhs: rhs, Expr: NewExpr()}
}

func (this BinaryOperator) GetLhs() interfaces.Expr {
	return this.Lhs
}

func (this BinaryOperator) GetRhs() interfaces.Expr {
	return this.Rhs
}
