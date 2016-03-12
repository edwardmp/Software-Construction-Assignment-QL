package stmt

import (
	"fmt"
	"ql/interfaces"
)

type If struct {
	Cond interfaces.Expr
	Body StmtList
	Stmt
}

func NewIf(condition interfaces.Expr, body StmtList, sourceInfo interface{}) If {
	return If{condition, body, NewStmt(sourceInfo)}
}

func NewIfNoSourceInfo(condition interfaces.Expr, body StmtList) If {
	return NewIf(condition, body, nil)
}

func (i If) String() string {
	return fmt.Sprintf("An if statement with condition %s and statement list %s", i.Cond, i.Body)
}

func (i If) EvalCondition() bool {
	return i.Cond.Eval(nil).(bool) // TODO symboltable
}
