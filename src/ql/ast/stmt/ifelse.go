package stmt

import "ql/interfaces"

type IfElse struct {
	condition     interfaces.Expr
	ifBody   StmtList
	elseBody StmtList
	Stmt
}

func NewIfElse(condition interfaces.Expr, ifBody StmtList, thenBody StmtList) IfElse {
	return IfElse{condition, ifBody, thenBody, NewStmt()}
}

func (this IfElse) Condition() interfaces.Expr {
	return this.condition
}

func (this IfElse) IfBody() interfaces.StmtList {
	return this.ifBody
}

func (this IfElse) ElseBody() interfaces.StmtList {
	return this.elseBody
}

func (this IfElse) EvalCondition(symbols interfaces.VarIdValueSymbols) bool {
	return this.condition.Eval(symbols).(bool)
}
