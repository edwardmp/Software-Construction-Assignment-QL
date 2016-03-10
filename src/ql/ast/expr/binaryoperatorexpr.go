package expr

import "ql/interfaces"

type BinaryOperatorExpr interface {
	interfaces.Expr
}

type BinaryOperator struct {
	Lhs, Rhs interfaces.Expr
	interfaces.Expr
}
