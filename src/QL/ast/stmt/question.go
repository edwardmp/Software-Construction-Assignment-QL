package stmt

import (
    "ql/ast/expr/lit"
	"ql/ast/vari"
	"ql/ast/visit"
)

type Question interface {
	GetLabel() lit.StrLit
	GetVarDecl() vari.VarDecl
	Accept(v visit.Visitor) interface{}
}
