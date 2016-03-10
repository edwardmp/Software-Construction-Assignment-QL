package stmt

import (
	"fmt"
	"ql/interfaces"
)

type IfElse struct {
	Cond     interfaces.Expr
	IfBody   StmtList
	ElseBody StmtList
}

func (i IfElse) String() string {
	return fmt.Sprintf("An if/else statement with condition %s", i.Cond)
}

func (i IfElse) EvalCondition() bool {
	return i.Cond.Eval(nil).(bool) // TODO symboltable
}
