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

func (ifElseStmt IfElse) String() string {
	return fmt.Sprintf("An if/else statement with condition %s", ifElseStmt.Cond)
}

func (ifElseStmt IfElse) EvalCondition() bool {
	return ifElseStmt.Cond.Eval().(bool)
}
