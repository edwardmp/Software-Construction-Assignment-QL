package expr

type StrLit struct {
	Value string
	Expr
}

func NewStrLit(value string) StrLit {
	return StrLit{Value: value, Expr: NewExpr()}
}

func (this StrLit) GetValue() string {
	return this.Value
}

// FIXME needed for test comparison due to strange behavior. Find better solution.
func (this StrLit) String() string {
	return this.Value
}
