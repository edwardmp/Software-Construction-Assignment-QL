package binaryoperatorexpr

import "ql/ast/expr"

type BinaryOperatorExpr interface {
	expr.Expr
}

type BinaryOperator struct {
	Lhs, Rhs expr.Expr
	BinaryOperatorExpr
}
