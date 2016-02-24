package stmt

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/visit"
)

type If struct {
	Cond expr.Expr
	Body StmtList
}

func (i If) Eval() {
	i.Cond.Eval(nil) // FIXME necessary?
	i.Body.Eval()
}

func (i If) String() string {
	return fmt.Sprintf("An if statement with condition %s and statement list %s", i.Cond, i.Body)
}

func (i If) EvalCondition() bool {
	return i.Cond.Eval(nil).(bool)
}

func (i If) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(i, s)
}
