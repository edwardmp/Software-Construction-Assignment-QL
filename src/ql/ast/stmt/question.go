package stmt

import (
	"ql/ast/vari"
	"ql/ast/visit"
)

type Question interface {
	GetLabelAsString() string
	GetVarDecl() vari.VarDecl
	Accept(v visit.Visitor, s interface{}) interface{}
}
