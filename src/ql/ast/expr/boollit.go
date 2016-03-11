package expr

import "fmt"

type BoolLit struct {
	Value bool
}

func NewBoolLit(value bool) BoolLit {
	return BoolLit{Value: value}
}

func (b BoolLit) GetValue() bool {
	return b.Value
}

func (b BoolLit) String() string {
	return fmt.Sprintf("%t", b.Value)
}
