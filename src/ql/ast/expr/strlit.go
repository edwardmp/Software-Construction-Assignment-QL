package expr

import "ql/interfaces"

type StrLit struct {
	stringValue StringValue
	Expr
}

func NewStrLit(value string) StrLit {
	return StrLit{stringValue: NewStringValue(value), Expr: NewExpr()}
}

func (this StrLit) Value() interfaces.Value {
	return this.stringValue
}
