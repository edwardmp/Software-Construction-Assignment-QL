package unaryoperatorexpr

import "ql/ast/expr"

type UnaryOperatorExpr interface {
	GetValue() expr.Expr
	expr.Expr
}
