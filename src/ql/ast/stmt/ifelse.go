package stmt

import (
	"fmt"
	"ql/ast/expr"
)

type IfElse struct {
	Cond     expr.Expr
	IfBody   StmtList
	ElseBody StmtList
}

func (i IfElse) String() string {
	return fmt.Sprintf("An if/else statement with condition %s", i.Cond)
}

func (i IfElse) EvalCondition() bool {
	return i.Cond.Eval(nil).(bool)
}
