package expr

import "ql/ast/visit"

type BinaryOperatorExpr interface {
	GetLhs() interface{}
	GetRhs() interface{}
	Accept(v visit.Visitor) interface{}
}
