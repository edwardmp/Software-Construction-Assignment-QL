package stmt

import (
	"ql/ast/expr/lit"
	"ql/ast/vari"
	"ql/ast/visit"
)

type Question interface {
	GetLabel() lit.StrLit
	GetLabelAsString() string
	GetVarDecl() vari.VarDecl
	Accept(v visit.Visitor, s interface{}) interface{}
}
