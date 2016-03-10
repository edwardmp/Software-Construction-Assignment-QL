package litexpr

import (
	"ql/ast/expr"
)

type LitExpr interface {
	expr.Expr
	String() string
}
