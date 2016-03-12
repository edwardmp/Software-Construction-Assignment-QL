package vari

import "ql/ast/node"

type Var struct {
	node.Node
}

func NewVar(sourceInfo interface{}) Var {
	return Var{node.NewNode(sourceInfo)}
}
