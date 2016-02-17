package unaryoperatorexpr

import "ql/ast/expr"

type UnaryOperatorExpr interface {
	GetValue() interface{}
    expr.Expr
}
