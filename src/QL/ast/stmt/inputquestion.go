package stmt

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/vari"
	"ql/ast/visit"
)

type InputQuestion struct {
	Label   expr.StrLit
	VarDecl vari.VarDecl
}

func (q InputQuestion) String() string {
	return fmt.Sprintf("A question with label %s, var decl %s", q.Label, q.VarDecl)
}

func (q InputQuestion) GetLabel() expr.StrLit {
	return q.Label
}

func (q InputQuestion) GetVarDecl() vari.VarDecl {
	return q.VarDecl
}

func (q InputQuestion) Accept(v visit.Visitor) interface{} {
	return v.Visit(q)
}
