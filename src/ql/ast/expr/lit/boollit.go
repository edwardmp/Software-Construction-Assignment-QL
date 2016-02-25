package lit

import (
	"fmt"
	"ql/ast/visit"
)

type BoolLit struct {
	Value bool
}

func (b BoolLit) GetValue() interface{} {
	return b.Value
}

func (b BoolLit) Eval(s interface{}) interface{} {
	return bool(b.Value)
}

func (b BoolLit) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(b, s)
}

func (b BoolLit) String() string {
	return fmt.Sprintf("%t", b)
}
