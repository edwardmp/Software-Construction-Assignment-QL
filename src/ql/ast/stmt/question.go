package stmt

import "ql/interfaces"

type Question struct {
	label   interfaces.StringLiteral
	varDecl interfaces.VarDecl
	Stmt
}

func NewQuestion(label interfaces.StringLiteral, varDecl interfaces.VarDecl) Question {
	return Question{label: label, varDecl: varDecl, Stmt: NewStmt()}
}

func (this Question) Label() interfaces.StringLiteral {
	return this.label
}

func (this Question) VarDecl() interfaces.VarDecl {
	return this.varDecl
}
