package litexpr

import (
	"fmt"
)

type BoolLit struct {
	Value bool
}

func (b BoolLit) Eval(s interface{}) interface{} {
	return bool(b.Value)
}

func (b BoolLit) String() string {
	return fmt.Sprintf("%t", b.Value)
}
