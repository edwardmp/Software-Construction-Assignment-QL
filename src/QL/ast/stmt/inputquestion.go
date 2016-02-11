package stmt

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/vari"
)

type InputQuestion struct {
	Label   expr.StrLit
	VarDecl vari.VarDecl
}

func (q InputQuestion) String() string {
	return fmt.Sprintf("A question with label %s, var decl %s", q.Label, q.VarDecl)
}

func (q InputQuestion) getLabel() expr.StrLit {
	return q.Label
}

func (q InputQuestion) getVarDecl() vari.VarDecl {
	return q.VarDecl
}
