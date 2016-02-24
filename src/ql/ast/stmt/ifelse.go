package stmt

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/visit"
)

type IfElse struct {
	Cond     expr.Expr
	IfBody   StmtList
	ElseBody StmtList
}

func (i IfElse) Eval() {
	i.Cond.Eval(nil) // FIXME necessary?
	i.IfBody.Eval()
	i.ElseBody.Eval()
}

func (i IfElse) String() string {
	return fmt.Sprintf("An if/else statement with condition %s", i.Cond)
}

func (i IfElse) EvalCondition() bool {
	return i.Cond.Eval(nil).(bool)
}

func (i IfElse) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(i, s)
}
