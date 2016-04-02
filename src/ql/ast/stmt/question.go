package stmt

import "ql/interfaces"

type Question struct {
	label   interfaces.StrLit
	varDecl interfaces.VarDecl
	Stmt
}

func NewQuestion(label interfaces.StrLit, varDecl interfaces.VarDecl) Question {
	return Question{label: label, varDecl: varDecl, Stmt: NewStmt()}
}

func (this Question) Label() interfaces.StrLit {
	return this.label
}

func (this Question) VarDecl() interfaces.VarDecl {
	return this.varDecl
}
