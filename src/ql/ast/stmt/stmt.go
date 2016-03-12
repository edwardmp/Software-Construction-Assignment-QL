package stmt

import "ql/ast/node"

type Stmt struct {
	node.Node
}

func NewStmt(sourceInfo interface{}) Stmt {
	return Stmt{node.NewNode(sourceInfo)}
}

func NewStmtNoSourceInfo() Stmt {
	return NewStmt(nil)
}
