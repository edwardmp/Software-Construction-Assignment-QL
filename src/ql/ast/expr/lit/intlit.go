package lit

import (
	"fmt"
	"ql/ast/visit"
)

type IntLit struct {
	Value int
}

func (i IntLit) GetValue() interface{} {
	return i.Value
}

func (i IntLit) Eval(s interface{}) interface{} {
	return i.Value
}

func (i IntLit) Accept(v visit.Visitor, s interface{}) interface{} {
	return v.Visit(i, s)
}

func (i IntLit) String() string {
	return fmt.Sprintf("%d", i.Value)
}
