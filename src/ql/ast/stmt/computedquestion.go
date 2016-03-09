package stmt

import (
	"fmt"
	"ql/ast/expr"
	"ql/ast/expr/litexpr"
	"ql/ast/vari"
)

type ComputedQuestion struct {
	Label       litexpr.StrLit
	VarDecl     vari.VarDecl
	Computation expr.Expr
}

func (c ComputedQuestion) String() string {
	return fmt.Sprintf("A question with label %s, var decl %s and computation", c.Label, c.VarDecl, c.Computation)
}

func (c ComputedQuestion) GetLabel() litexpr.StrLit {
	return c.Label
}

func (c ComputedQuestion) GetLabelAsString() string {
	return c.Label.Value
}

func (c ComputedQuestion) GetVarDecl() vari.VarDecl {
	return c.VarDecl
}