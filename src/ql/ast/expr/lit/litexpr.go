package lit

import (
	"fmt"
	"ql/ast/expr"
)

type Lit interface {
	GetValue() interface{}
	expr.Expr
	fmt.Stringer
}
