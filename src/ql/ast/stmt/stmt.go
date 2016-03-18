package stmt

import (
	"fmt"
	"ql/ast/node"
)

type Stmt struct {
	node.Node
}

func NewStmt(sourceInfo interface{}) Stmt {
	return Stmt{node.NewNode(sourceInfo)}
}

func NewStmtNoSourceInfo() Stmt {
	return NewStmt(nil)
}

// slices don't support equality checking, so have to do it like this, by printing its value
func SlicesEqual(a StmtList, b StmtList) bool {
	if fmt.Sprintf("%v", a) != fmt.Sprintf("%v", b) {
		return false
	}

	return true
}
