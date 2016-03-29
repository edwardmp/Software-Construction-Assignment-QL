package expr

type ValueType struct {
	Expr
	TypeString string
}

func NewValueType(typeString string) ValueType {
	return ValueType{NewExpr(), typeString}
}

func (this ValueType) String() string {
	return this.TypeString
}
