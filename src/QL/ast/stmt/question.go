package stmt

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/vari"
)

type Question struct {
	Label   expr.StrLit
	VarDecl vari.VarDecl
}

func (q Question) String() string {
	return fmt.Sprintf("A question with label %s, var decl %s", q.Label, q.VarDecl)
}
