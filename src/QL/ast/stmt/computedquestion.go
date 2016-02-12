package stmt

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/vari"
	"ql/ast/visit"
)

type ComputedQuestion struct {
	Label       expr.StrLit
	VarDecl     vari.VarDecl
	Computation expr.Expr
}

func (c ComputedQuestion) String() string {
	return fmt.Sprintf("A question with label %s, var decl %s and computation", c.Label, c.VarDecl, c.Computation)
}

func (c ComputedQuestion) GetLabel() expr.StrLit {
	return c.Label
}

func (c ComputedQuestion) GetVarDecl() vari.VarDecl {
	return c.VarDecl
}

func (c ComputedQuestion) Accept(v visit.Visitor) interface{} {
	return v.Visit(c)
}
