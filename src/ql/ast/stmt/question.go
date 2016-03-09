package stmt

import (
	"ql/ast/expr/litexpr"
	"ql/ast/vari"
)

type Question interface {
	GetLabel() litexpr.StrLit
	GetLabelAsString() string
	GetVarDecl() vari.VarDecl
}
