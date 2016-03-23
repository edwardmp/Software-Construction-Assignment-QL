package stmt

import (
	"fmt"
	"ql/ast/node"
	"ql/interfaces"
)

type Stmt struct {
	node.Node
}

func NewStmt() Stmt {
	return Stmt{node.NewNode()}
}

// slices don't support equality checking, so have to do it like this, by printing its value
func SlicesEqual(a, b interfaces.StmtList) bool {
	if fmt.Sprintf("%v", a) != fmt.Sprintf("%v", b) {
		return false
	}

	return true
}
