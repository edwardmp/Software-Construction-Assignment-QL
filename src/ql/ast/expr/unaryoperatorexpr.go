package expr

import "ql/interfaces"

type UnaryOperatorExpr interface {
	interfaces.Expr
}

type UnaryOperator struct {
	Value interfaces.Expr
	interfaces.Expr
}
