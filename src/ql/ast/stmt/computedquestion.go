package stmt

import "ql/interfaces"

type ComputedQuestion struct {
    label       interfaces.StrLit
	varDecl     interfaces.VarDecl
	computation interfaces.Expr
	Stmt
}

func NewComputedQuestion(label interfaces.StrLit, varDecl interfaces.VarDecl, computation interfaces.Expr) ComputedQuestion {
	return ComputedQuestion{label, varDecl, computation, NewStmt()}
}

func (this ComputedQuestion) Label() interfaces.StrLit {
	return this.label
}

func (this ComputedQuestion) Computation() interfaces.Expr {
	return this.computation
}

func (this ComputedQuestion) VarDecl() interfaces.VarDecl {
	return this.varDecl
}
