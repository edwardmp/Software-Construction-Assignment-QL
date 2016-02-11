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

func (q ComputedQuestion) String() string {
	return fmt.Sprintf("A question with label %s, var decl %s and computation", q.Label, q.VarDecl, q.Computation)
}

func (q ComputedQuestion) GetLabel() expr.StrLit {
	return q.Label
}

func (q ComputedQuestion) GetVarDecl() vari.VarDecl {
	return q.VarDecl
}

func (q ComputedQuestion) Accept(v visit.Visitor) interface{} {
	return v.Visit(q)
}
