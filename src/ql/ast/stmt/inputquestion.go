package stmt

import (
	"fmt"
	"ql/ast/expr/lit"
	"ql/ast/vari"
	"ql/ast/visit"
)

type InputQuestion struct {
	Label   lit.StrLit
	VarDecl vari.VarDecl
}

func (i InputQuestion) String() string {
	return fmt.Sprintf("A question with label %s, var decl %s", i.Label, i.VarDecl)
}

func (i InputQuestion) GetLabelAsString() string {
	return i.Label.GetValue().(string)
}

func (i InputQuestion) GetVarDecl() vari.VarDecl {
	return i.VarDecl
}

func (i InputQuestion) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(i, s)
}
