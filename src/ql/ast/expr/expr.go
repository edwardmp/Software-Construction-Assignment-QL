package expr

import (
	"ql/ast/visit"
)

type Expr interface {
	Eval() interface{}
	Accept(v visit.Visitor, s interface{}) interface{}
}
