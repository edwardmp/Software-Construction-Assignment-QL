package stmt

import "QL/ast/expr"

type IfElse struct {
	Cond     expr.Expr
	IfBody   Stmt
	ElseBody Stmt
}
