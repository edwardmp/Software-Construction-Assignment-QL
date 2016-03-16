package stmt

import (
	"fmt"
	"ql/interfaces"
)

type IfElse struct {
	Cond     interfaces.Expr
	IfBody   StmtList
	ElseBody StmtList
	Stmt
}

func NewIfElse(condition interfaces.Expr, ifBody StmtList, thenBody StmtList, sourceInfo interface{}) IfElse {
	return IfElse{condition, ifBody, thenBody, NewStmt(sourceInfo)}
}

func NewIfElseNoSourceInfo(condition interfaces.Expr, ifBody StmtList, thenBody StmtList) IfElse {
	return NewIfElse(condition, ifBody, thenBody, nil)
}

func (i IfElse) String() string {
	return fmt.Sprintf("An if/else statement with condition %s", i.Cond)
}

func (i IfElse) EvalCondition() bool {
	return i.Cond.Eval(nil).(bool) // TODO symboltable
}
