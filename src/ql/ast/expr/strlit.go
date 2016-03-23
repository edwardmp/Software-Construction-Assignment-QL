package expr

type StrLit struct {
	value string
	Expr
}

func NewStrLit(value string) StrLit {
	return StrLit{value: value, Expr: NewExpr()}
}

func (this StrLit) Value() string {
	return this.value
}

// FIXME needed for test comparison due to strange behavior. Find better solution.
func (this StrLit) String() string {
	return this.value
}
