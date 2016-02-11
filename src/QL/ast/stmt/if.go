package stmt

import (
	"fmt"
	"ql/ast/expr"
)

type If struct {
	Cond expr.Expr
	Body StmtList
}

func (ifStmt If) String() string {
	return fmt.Sprintf("An if statement with condition %s", ifStmt.Cond)
}

func (ifStmt If) EvalCondition() bool {
	return ifStmt.Cond.Eval().(bool)
}
