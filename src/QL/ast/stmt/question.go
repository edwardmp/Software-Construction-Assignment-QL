package stmt

import (
	"ql/ast/expr"
	"ql/ast/vari"
)

type Question interface {
	getLabel() expr.StrLit
	getVarDecl() vari.VarDecl
}
