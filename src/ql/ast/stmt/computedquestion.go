package stmt

import (
	"fmt"
	"ql/interfaces"
)

type ComputedQuestion struct {
	Label       interfaces.StrLit
	VarDecl     interfaces.VarDecl
	Computation interfaces.Expr
}

func (c ComputedQuestion) String() string {
	return fmt.Sprintf("A question with label %s, var decl %s and computation", c.Label, c.VarDecl, c.Computation)
}

func (c ComputedQuestion) GetLabel() interfaces.StrLit {
	return c.Label
}

func (c ComputedQuestion) GetLabelAsString() string {
	return c.Label.GetValue()
}

func (c ComputedQuestion) GetComputation() interfaces.Expr {
	return c.Computation
}

func (c ComputedQuestion) GetVarDecl() interfaces.VarDecl {
	return c.VarDecl
}
