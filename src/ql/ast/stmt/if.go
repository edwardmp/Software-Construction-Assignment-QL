package stmt

import "ql/interfaces"

type If struct {
	condition interfaces.Expr
	body StmtList
	Stmt
}

func NewIf(condition interfaces.Expr, body StmtList) If {
	return If{condition, body, NewStmt()}
}

func (this If) Body() interfaces.StmtList {
	return this.body
}

func (this If) Condition() interfaces.Expr {
	return this.condition
}

func (this If) EvalCondition(symbolTable interfaces.VarIdValueSymbols) bool {
	return this.condition.Eval(symbolTable).(bool)
}
