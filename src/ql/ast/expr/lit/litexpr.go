package lit

import "ql/ast/expr"

type Lit interface {
	GetValue() interface{}
	expr.Expr
}
