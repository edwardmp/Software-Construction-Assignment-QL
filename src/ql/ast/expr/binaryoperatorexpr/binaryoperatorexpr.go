package binaryoperatorexpr

import "ql/ast/expr"

type BinaryOperatorExpr interface {
	GetLhs() expr.Expr
	GetRhs() expr.Expr
	expr.Expr
}
