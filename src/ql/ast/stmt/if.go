package stmt

import "ql/interfaces"

type If struct {
	Cond interfaces.Expr
	Body StmtList
	Stmt
}

func NewIf(condition interfaces.Expr, body StmtList, sourceInfo interface{}) If {
	return If{condition, body, NewStmt(sourceInfo)}
}

func NewIfNoSourceInfo(condition interfaces.Expr, body StmtList) If {
	return NewIf(condition, body, nil)
}

func (this If) GetBody() interfaces.StmtList {
	return this.Body
}

func (this If) GetCondition() interfaces.Expr {
	return this.Cond
}

func (this If) EvalCondition(symbolTable interfaces.VarIdValueSymbols) bool {
	return this.Cond.Eval(symbolTable).(bool)
}
