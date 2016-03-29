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
