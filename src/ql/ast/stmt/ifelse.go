package stmt

import "ql/interfaces"

type IfElse struct {
	Cond     interfaces.Expr
	IfBody   StmtList
	ElseBody StmtList
	Stmt
}

func NewIfElse(condition interfaces.Expr, ifBody StmtList, thenBody StmtList) IfElse {
	return IfElse{condition, ifBody, thenBody, NewStmt()}
}

func (this IfElse) GetCondition() interfaces.Expr {
	return this.Cond
}

func (this IfElse) GetIfBody() interfaces.StmtList {
	return this.IfBody
}

func (this IfElse) GetElseBody() interfaces.StmtList {
	return this.ElseBody
}

func (this IfElse) EvalCondition(symbols interfaces.VarIdValueSymbols) bool {
	return this.Cond.Eval(symbols).(bool)
}
