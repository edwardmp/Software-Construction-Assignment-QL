package unaryoperatorexpr

import "ql/ast/expr"

type UnaryOperatorExpr interface {
	expr.Expr
}

type UnaryOperator struct {
	Value expr.Expr
	UnaryOperatorExpr
}
