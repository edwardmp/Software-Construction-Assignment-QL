package stmt

import (
	"fmt"
	"ql/interfaces"
)

type If struct {
	Cond interfaces.Expr
	Body StmtList
}

func (i If) String() string {
	return fmt.Sprintf("An if statement with condition %s and statement list %s", i.Cond, i.Body)
}

func (i If) EvalCondition() bool {
	return i.Cond.Eval(nil).(bool)
}
