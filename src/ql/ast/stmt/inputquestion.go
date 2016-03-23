package stmt

import "ql/interfaces"

type InputQuestion struct {
	label   interfaces.StrLit
	varDecl interfaces.VarDecl
	Stmt
}

func NewInputQuestion(label interfaces.StrLit, varDecl interfaces.VarDecl) InputQuestion {
	return InputQuestion{label, varDecl, NewStmt()}
}

func (this InputQuestion) Label() interfaces.StrLit {
	return this.label
}

func (this InputQuestion) VarDecl() interfaces.VarDecl {
	return this.varDecl
}
