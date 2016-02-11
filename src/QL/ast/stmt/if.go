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

func (i If) String() string {
	return fmt.Sprintf("An if statement with condition %s and statement list %s", i.Cond, i.Body)
}

func (i If) EvalCondition() bool {
	return i.Cond.Eval().(bool)
}

func (i If) Accept(v visit.Visitor) interface{} {
	return v.Visit(i)
}
