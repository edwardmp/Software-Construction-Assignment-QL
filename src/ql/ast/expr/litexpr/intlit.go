package litexpr

import (
	"fmt"
)

type IntLit struct {
	Value int
}

func (i IntLit) Eval(s interface{}) interface{} {
	return i.Value
}

func (i IntLit) String() string {
	return fmt.Sprintf("%d", i.Value)
}
