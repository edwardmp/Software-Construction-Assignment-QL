package stmt

import (
	"ql/ast/expr/litexpr"
	"ql/ast/vari"
	"ql/ast/visit"
)

type Question interface {
	GetLabel() litexpr.StrLit
	GetLabelAsString() string
	GetVarDecl() vari.VarDecl
	Accept(v visit.Visitor, s interface{}) interface{}
}
