package expr

import (
	"fmt"
)

type IntLit struct {
	Value int
}

func NewIntLit(value int) IntLit {
	return IntLit{Value: value}
}

func (i IntLit) Eval(s interface{}) interface{} {
	return i.Value
}

func (i IntLit) String() string {
	return fmt.Sprintf("%d", i.Value)
}
