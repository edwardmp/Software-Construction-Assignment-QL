package stmt

import "ql/ast/visit"

type Conditional interface {
	EvalCondition() bool
	Accept(v visit.Visitor) interface{}
}
