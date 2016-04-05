package stmt

import "ql/interfaces"

type IfElse struct {
	condition interfaces.Expr
	ifBody    StmtList
	elseBody  StmtList
	Stmt
}

func NewIfElse(condition interfaces.Expr, ifBody StmtList, elseBody StmtList) IfElse {
	return IfElse{condition: condition, ifBody: ifBody, elseBody: elseBody, Stmt: NewStmt()}
}

func (this IfElse) IfBody() interfaces.StmtList {
	return this.ifBody
}

func (this IfElse) ElseBody() interfaces.StmtList {
	return this.elseBody
}

func (this IfElse) Condition() interfaces.Expr {
	return this.condition
}

func (this IfElse) EvalCondition(symbols interfaces.VarIdValueSymbols) interfaces.BoolValue {
	return this.condition.Eval(symbols).(interfaces.BoolValue)
}
