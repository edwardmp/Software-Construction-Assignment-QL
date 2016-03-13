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

func (this If) GetBody() interfaces.StmtList {
	return this.Body
}

func (this If) GetCondition() interfaces.Expr {
	return this.Cond
}

func (this If) String() string {
	return fmt.Sprintf("An if statement with condition %s and statement list %s", this.Cond, this.Body)
}

func (this If) EvalCondition(symbolTable interfaces.Symbols) bool {
	return this.Cond.Eval(symbolTable).(bool)
}
