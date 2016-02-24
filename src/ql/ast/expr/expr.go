package expr

import (
	"ql/ast/visit"
)

type Expr interface {
	Eval(s interface{}) interface{}
	Accept(v visit.Visitor, s interface{}) interface{}
}
