package stmt

import "QL/ast/expr"

type If struct {
	Cond expr.Expr
	Body Stmt
}
