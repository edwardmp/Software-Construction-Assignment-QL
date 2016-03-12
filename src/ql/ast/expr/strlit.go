package expr

type StrLit struct {
	Value string
	Expr
}

func NewStrLit(value string, sourceInfo interface{}) StrLit {
	return StrLit{Value: value, Expr: NewExpr(sourceInfo)}
}

func NewStrLitNoSourceInfo(value string) StrLit {
	return NewStrLit(value, nil)
}

func (s StrLit) GetValue() string {
	return s.Value
}

func (s StrLit) String() string {
	return s.Value
}
