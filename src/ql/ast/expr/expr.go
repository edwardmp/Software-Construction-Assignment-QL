package expr

import (
	"ql/ast/node"
)

type Expr struct {
	node.Node
}

func NewExpr(sourceInfo interface{}) Expr {
	return Expr{node.NewNode(sourceInfo)}
}

func NewExprNoSourceInfo() Expr {
	return NewExpr(nil)
}
