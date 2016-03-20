package stmt

import (
	"fmt"
	"ql/ast/node"
)

type Stmt struct {
	node.Node
}

func NewStmt() Stmt {
	return Stmt{node.NewNode()}
}

// slices don't support equality checking, so have to do it like this, by printing its value
func SlicesEqual(a, b StmtList) bool {
	if fmt.Sprintf("%v", a) != fmt.Sprintf("%v", b) {
		return false
	}

	return true
}
