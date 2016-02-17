package ast

import "ql/ast/visit"

type Node interface {
	Accept(v visit.Visitor) interface{}
}
