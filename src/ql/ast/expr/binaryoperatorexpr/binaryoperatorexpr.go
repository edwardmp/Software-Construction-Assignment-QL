package binaryoperatorexpr

import "ql/ast/expr"

type BinaryOperatorExpr interface {
	GetLhs() interface{}
	GetRhs() interface{}
	expr.Expr
}
