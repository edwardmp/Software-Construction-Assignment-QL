package typechecker

import "ql/ast/visit"

type TypeChecker struct {
	visit.Visitor
	ErrorsEncountered []error
}
