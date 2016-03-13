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

func (i IfElse) GetCondition() interfaces.Expr {
	return i.Cond
}

func (i IfElse) GetIfBody() interfaces.StmtList {
	return i.IfBody
}

func (i IfElse) GetElseBody() interfaces.StmtList {
	return i.ElseBody
}

func (i IfElse) String() string {
	return fmt.Sprintf("An if/else statement with condition %s", i.Cond)
}

func (i IfElse) EvalCondition(symbolTable interfaces.SymbolTable) bool {
	return i.Cond.Eval(symbolTable).(bool)
}
