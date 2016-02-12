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

func (i InputQuestion) String() string {
	return fmt.Sprintf("A question with label %s, var decl %s", i.Label, i.VarDecl)
}

func (i InputQuestion) GetLabel() expr.StrLit {
	return i.Label
}

func (i InputQuestion) GetVarDecl() vari.VarDecl {
	return i.VarDecl
}

func (i InputQuestion) Accept(v visit.Visitor) interface{} {
	return v.Visit(i)
}
