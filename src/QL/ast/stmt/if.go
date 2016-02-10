package stmt

import "QL/ast/expr"

If struct {
	Cond expr
	Body Stmt
}