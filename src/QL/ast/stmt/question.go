package stmt

import (
	"ql/ast/expr"
	"ql/ast/vari"
	"ql/ast/visit"
)

type Question interface {
	GetLabel() expr.StrLit
	GetVarDecl() vari.VarDecl
	Accept(v visit.Visitor) interface{}
}
