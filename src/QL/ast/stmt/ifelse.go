package stmt

import "QL/ast/expr"

IfElse struct {
	Cond expr
	IfBody Stmt
	ElseBody Stmt
}