package vari

import "ql/ast/node"

type Var struct {
	node.Node
}

func NewVar() Var {
	return Var{node.NewNode()}
}
