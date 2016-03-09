package litexpr

import (
	"ql/ast/expr"
)

type Lit interface {
	expr.Expr
	String() string
}
