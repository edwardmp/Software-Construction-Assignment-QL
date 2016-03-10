package expr

import "ql/interfaces"

type BinaryOperator struct {
	Lhs, Rhs interfaces.Expr
}

func (b BinaryOperator) GetLhs() interfaces.Expr {
	return b.Lhs
}

func (b BinaryOperator) GetRhs() interfaces.Expr {
	return b.Rhs
}
