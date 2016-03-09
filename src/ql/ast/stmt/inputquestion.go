package stmt

import (
	"fmt"
	"ql/ast/expr/litexpr"
	"ql/ast/vari"
)

type InputQuestion struct {
	Label   litexpr.StrLit
	VarDecl vari.VarDecl
}

func (i InputQuestion) String() string {
	return fmt.Sprintf("A question with label %s, var decl %s", i.Label, i.VarDecl)
}

func (i InputQuestion) GetLabel() litexpr.StrLit {
	return i.Label
}

func (i InputQuestion) GetLabelAsString() string {
	return i.Label.Value
}

func (i InputQuestion) GetVarDecl() vari.VarDecl {
	return i.VarDecl
}